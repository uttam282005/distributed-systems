# Distributed Systems Theory Repository

A rigorous, graduate-level reference repository and self-study curriculum dedicated to the mathematical foundations, algorithm specifications, correctness invariants, and architectural trade-offs of distributed systems.

---

## 1. Repository Philosophy and Scope

This repository focuses strictly on **theoretical and architectural principles**. It intentionally avoids runnable implementation exercises, benchmarks, framework-specific libraries, or disposable codebases.

Every topic is developed with:
- **Precise System Models**: Explicit specifications of processes, link models, timing assumptions, and failure modes.
- **Mathematical Formulations**: Rigorous definitions of invariants, partial orders, execution histories, and equivalence classes.
- **Language-Independent Pseudocode**: Clean, formal event-driven specifications detailing process state, message schemas, and handlers.
- **Execution Traces & Counterexamples**: Concrete step-by-step traces showing both correct executions and failure scenarios.
- **Complexity Characterizations**: Exact round complexity, message counts, and space overheads.
- **Authoritative Citations**: Direct links and citations to seminal computer science literature.

---

## 2. Repository Navigation

- **[Mastery Roadmap (`ROADMAP.md`)](./ROADMAP.md)**: A dependency-aware 7-stage learning journey with objectives, self-check questions, and mastery criteria.
- **[Formal Glossary (`GLOSSARY.md`)](./GLOSSARY.md)**: Precise definitions of key concepts, models, and protocols.
- **[Bibliography & References (`references/README.md`)](./references/README.md)**: Primary textbooks, seminal papers, and lecture notes.
- **[Curriculum Progress (`PROGRESS.md`)](./PROGRESS.md)**: Real-time status tracker of all modules and substantive documentation.

---

## 3. Curriculum Structure

### [00. Foundations](./00-foundations/)
1. [`distributed-systems-overview.md`](./00-foundations/distributed-systems-overview.md) — What defines a distributed system, partial failures, concurrency without shared memory.
2. [`system-models.md`](./00-foundations/system-models.md) — Formalizing processes, networks, timing, and adversaries.
3. [`synchrony-and-asynchrony.md`](./00-foundations/synchrony-and-asynchrony.md) — Synchronous ($\mathcal{SYNCH}$), asynchronous ($\mathcal{ASYNCH}$), and partially synchronous ($\mathcal{PARTIAL}$) systems; Global Stabilization Time (GST).
4. [`failure-models.md`](./00-foundations/failure-models.md) — Crash-stop, crash-recovery (stable storage), and Byzantine (fail-arbitrary) failure modes.
5. [`safety-and-liveness.md`](./00-foundations/safety-and-liveness.md) — Alpern-Schneider trace characterization, finite prefix violation vs infinite horizon guarantees.
6. [`impossibility-results.md`](./00-foundations/impossibility-results.md) — Two Generals, FLP impossibility, CAP theorem, and lower bounds.

### [01. Time and Ordering](./01-time-and-ordering/)
1. [`physical-clocks.md`](./01-time-and-ordering/physical-clocks.md) — Quartz drift, skew, NTP Marzullo bounds, slewing vs stepping, and Spanner TrueTime uncertainty intervals.
2. [`happens-before.md`](./01-time-and-ordering/happens-before.md) — Lamport's strict partial order ($\to$), concurrency ($\parallel$), causal pasts, and relativistic light cones.
3. [`lamport-clocks.md`](./01-time-and-ordering/lamport-clocks.md) — Scalar logical time, Clock Condition, total order extensions, and tie-breaking.
4. [`vector-clocks.md`](./01-time-and-ordering/vector-clocks.md) — Vector timestamping, causal history isomorphism, vector lattice ordering, and concurrency detection.

### [02. Communication and Coordination](./02-communication-and-coordination/)
1. [`message-passing.md`](./02-communication-and-coordination/message-passing.md) — Unicast primitives, FIFO channels, and link reliability layers.
2. [`rpc.md`](./02-communication-and-coordination/rpc.md) — Remote Procedure Call mechanics, stub marshalling, and delivery semantics (at-most-once, at-least-once, exactly-once).
3. [`reliable-broadcast.md`](./02-communication-and-coordination/reliable-broadcast.md) — Best-effort, Reliable, FIFO, Causal, and Total Order (Atomic) broadcast protocols and their lattice hierarchy.
4. [`failure-detectors.md`](./02-communication-and-coordination/failure-detectors.md) — Chandra-Toueg axiomatic detector classes ($\mathcal{P}, \diamondsuit\mathcal{P}, \diamondsuit\mathcal{S}, \Omega$), completeness, and accuracy.
5. [`leader-election.md`](./02-communication-and-coordination/leader-election.md) — Symmetry breaking, Bully algorithm, ring election, and term-based fencing.
6. [`quorums.md`](./02-communication-and-coordination/quorums.md) — Formal quorum systems, majority quorums, read/write overlap invariants ($r+w > n$), and Byzantine quorum intersections ($2f+1$).

### [03. Replication and Consistency](./03-replication-and-consistency/)
1. [`replication-fundamentals.md`](./03-replication-and-consistency/replication-fundamentals.md) — Active vs passive replication, state-based vs operation-based transfer.
2. [`primary-backup.md`](./03-replication-and-consistency/primary-backup.md) — Primary-backup topologies, failover hazards, split-brain, and fencing tokens.
3. [`shared-registers.md`](./03-replication-and-consistency/shared-registers.md) — Lamport register hierarchy: Safe, Regular, and Atomic shared registers.
4. [`abd-algorithm.md`](./03-replication-and-consistency/abd-algorithm.md) — Attiya-Bar-Noy-Dolev algorithm: SWMR & MWMR protocols, quorum queries, and why read write-back is necessary.
5. [`linearizability.md`](./03-replication-and-consistency/linearizability.md) — Herlihy-Wing formulation, linearization points, real-time precedence order, and composability.
6. [`sequential-consistency.md`](./03-replication-and-consistency/sequential-consistency.md) — Lamport sequential consistency, program order preservation, and non-composability.
7. [`eventual-consistency.md`](./03-replication-and-consistency/eventual-consistency.md) — Optimistic replication, conflict resolution (LWW vs Multi-value), and CRDTs.
8. [`state-machine-replication.md`](./03-replication-and-consistency/state-machine-replication.md) — Deterministic state machines, log-driven transitions, and equivalence to total order broadcast.

### [04. Consensus](./04-consensus/)
1. [`consensus-problem.md`](./04-consensus/consensus-problem.md) — Formal problem statement: Agreement, Validity, Termination.
2. [`flp-impossibility.md`](./04-consensus/flp-impossibility.md) — Fischer-Lynch-Paterson proof deconstruction: bivalent initial states, commutativity, and critical step preservation.
3. [`paxos.md`](./04-consensus/paxos.md) — Single-decree Synod, Phase 1 (Prepare/Promise), Phase 2 (Accept/Accepted), and Multi-Paxos.
4. [`raft.md`](./04-consensus/raft.md) — Leader election, log replication, commit rules, leader completeness invariant, and cluster membership changes.
5. [`byzantine-fault-tolerance.md`](./04-consensus/byzantine-fault-tolerance.md) — Byzantine generals, PBFT (Pre-Prepare, Prepare, Commit), and the $3f+1$ lower bound.

### [05. Storage and Transactions](./05-storage-and-transactions/)
1. [`partitioning-and-sharding.md`](./05-storage-and-transactions/partitioning-and-sharding.md) — Horizontal partitioning schemes, range vs hash partitioning, cross-shard queries.
2. [`consistent-hashing.md`](./05-storage-and-transactions/consistent-hashing.md) — Hash ring topology, virtual nodes, load balancing, and $O(K/N)$ key reassignment.
3. [`replicated-logs.md`](./05-storage-and-transactions/replicated-logs.md) — Write-Ahead Logging (WAL), segment indexing, LSM-trees, and log compaction.
4. [`two-phase-commit.md`](./05-storage-and-transactions/two-phase-commit.md) — Atomic commitment, prepare and commit phases, coordinator failure, in-doubt states, and Paxos Commit.
5. [`distributed-transactions.md`](./05-storage-and-transactions/distributed-transactions.md) — Distributed ACID, concurrency control (2PL, OCC, SSI), and cross-shard atomicity.
6. [`mvcc.md`](./05-storage-and-transactions/mvcc.md) — Multi-Version Concurrency Control, snapshot isolation, write skew, and Spanner commit-wait.
7. [`recovery-and-rebalancing.md`](./05-storage-and-transactions/recovery-and-rebalancing.md) — Crash recovery protocols, checkpointing, replica rebuilds, and state rebalancing.

### [06. Production Systems](./06-production-systems/)
1. [`timeouts-and-retries.md`](./06-production-systems/timeouts-and-retries.md) — Uncorrelated retry storms, exponential backoff with full jitter, deadline propagation.
2. [`idempotency.md`](./06-production-systems/idempotency.md) — Mathematical idempotency ($f(f(x)) = f(x)$), deduplication keys, and tombstoning.
3. [`backpressure.md`](./06-production-systems/backpressure.md) — Queue overflows, proactive flow control, reactive rate limiting, and token buckets.
4. [`load-balancing.md`](./06-production-systems/load-balancing.md) — Layer 4 vs Layer 7 routing, power-of-two-choices, health check stampedes.
5. [`service-discovery.md`](./06-production-systems/service-discovery.md) — Dynamic registration, heartbeating, consensus-backed discovery, and DNS limitations.
6. [`observability.md`](./06-production-systems/observability.md) — Distributed tracing, causality propagation (W3C Trace Context), metrics, and tail-latency profiling.
7. [`disaster-recovery.md`](./06-production-systems/disaster-recovery.md) — Multi-region replication topologies, RPO/RTO invariants, split-brain avoidance during data center failover.

---

## 4. Standard Note Convention

Every note follows a structured theoretical template:
1. **Problem Statement & Motivation**: Core challenges in distributed execution.
2. **System Model**: Mathematical assumptions regarding processes, channels, synchrony, and adversaries.
3. **Formal Specification**: Invariants, safety, and liveness definitions.
4. **Algorithm Specification**: Clean, language-independent event-driven pseudocode.
5. **Execution Traces**: Concrete step-by-step numbers and state transitions.
6. **Correctness Proofs**: Invariant proofs and contradiction sketches.
7. **Failure Scenarios**: Behavior under crashes, latency spikes, and network partitions.
8. **Complexity Analysis**: Round and message complexity.
9. **Authoritative References**: Citations to foundational literature.
