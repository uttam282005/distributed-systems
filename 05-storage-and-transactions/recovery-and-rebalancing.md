# Recovery, Checkpointing, and Cluster Rebalancing

Crash recovery and cluster rebalancing ensure that distributed nodes restore consistent state following unexpected outages and maintain balanced data distributions as nodes dynamically join or depart.

---

## 1. Crash Recovery and Stable Storage

In the crash-recovery model, a server may crash at any time, instantly wiping all volatile state (RAM, active connection sockets, OS buffers).

### The Three Phases of ARIES Recovery [Mohan et al. 1992]
When a crashed node reboots, it inspects its non-volatile Write-Ahead Log (WAL) starting from the most recent checkpoint:

```mermaid
flowchart LR
    CP["Checkpoint Record"] --> P1["1. Analysis Phase<br/>Scans WAL forward<br/>Reconstructs Dirty Page Table & Active Transactions"]
    P1 --> P2["2. Redo Phase<br/>Replays all logged changes forward<br/>Restores DB to exact state at crash instant"]
    P2 --> P3["3. Undo Phase<br/>Scans WAL backward<br/>Rolls back all uncommitted active transactions"]
```

---

## 2. Checkpointing Techniques

Without checkpointing, crash recovery requires replaying the entire history of the log from the beginning of time ($t = 0$). Checkpointing bounds recovery duration by periodically writing an in-memory state snapshot to disk and truncating the preceding log.

### 2.1 Consistent Global Checkpointing: Chandy-Lamport Algorithm [1985]
In a distributed system, taking a global snapshot across multiple processes cannot stop the world.
The **Chandy-Lamport Algorithm** takes a causally consistent distributed snapshot using special **Marker Messages**:

```mermaid
sequenceDiagram
    autonumber
    participant P1 as Process 1 (Initiator)
    participant P2 as Process 2

    Note over P1: P1 records local state S1
    P1->>P2: Marker M (along channel C12)
    Note over P1: P1 starts recording incoming messages on incoming channels...
    
    Note over P2: P2 receives Marker M for first time!
    Note over P2: P2 immediately records local state S2
    P2->>P1: Marker M (along channel C21)
    Note over P2: P2 records channel state...
    
    Note over P1: P1 receives Marker from P2: Stops recording channel C21!
    Note over P1, P2: Consistent Distributed Snapshot S1 + S2 + in-flight messages completed!
```

---

## 3. Cluster Rebalancing and Partition Migration

When cluster membership changes (adding a new server to scale capacity or removing a decommissioned machine), the cluster must migrate a subset of partitions to the new node without going offline.

### 3.1 Migration Workflow

```mermaid
sequenceDiagram
    autonumber
    participant NewNode as Target Node (New)
    participant Source as Source Node (Active)
    participant Client as Client

    Note over Source: Source actively serves reads/writes for Shard 4
    Client->>Source: write(key, val) -> OK
    
    Note over Source, NewNode: 1. Stream Base Snapshot to New Node
    Source-)NewNode: Stream base SSTables / snapshot files
    
    Note over Source, NewNode: 2. Stream Incremental Write Delta
    Source-)NewNode: Stream WAL mutations accumulated during snapshot transfer
    
    Note over Source, NewNode: 3. Cutover Coordination
    Note over Source: Brief pause (< 10ms) to drain in-flight requests
    Source->>NewNode: Hand over primary ownership of Shard 4
    Note over NewNode: New Node publishes routing update to Cluster Metadata
    
    Client->>NewNode: write(key, val) -> OK
```

### 3.2 Operational Safeguards
1. **Throttled Bandwidth**: Rebalancing network streams must be throttled (e.g., max 50 MB/s per node) to prevent background sync from saturating switches and causing production query timeouts.
2. **Dual-Writing during Cutover**: Replicas maintain dual buffers or proxy writes to the target shard during final cutover.

---

## 4. Revision Questions

1. Describe the Analysis, Redo, and Undo phases of the ARIES recovery protocol.
2. How does the Chandy-Lamport algorithm record in-flight messages in transit along communication channels?
3. What is the danger of unthrottled partition rebalancing in a production distributed database?
4. How do virtual nodes (vnodes) parallelize data transfer during node addition compared to single-token hashing?

---

## 5. Authoritative References

- **[Mohan et al. 1992]** C. Mohan et al. *ARIES: A Transaction Recovery Method Supporting Fine-Granularity Locking and Partial Rollbacks Using Write-Ahead Logging*. ACM TODS, 1992.
- **[Chandy, Lamport 1985]** K. Mani Chandy and Leslie Lamport. *Distributed Snapshots: Determining Global States of Distributed Systems*. ACM TOCS, 3(1):63–75, 1985.
- **[Kleppmann 2017]** Martin Kleppmann. *Designing Data-Intensive Applications*, Chapter 6 & 11. O'Reilly, 2017.
