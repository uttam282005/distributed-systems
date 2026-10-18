# Diagrams Index

This directory provides an architectural index of the Mermaid sequence diagrams, state machines, and system topologies embedded throughout the theoretical notes.

All diagrams are rendered natively using GitHub Flavored Markdown Mermaid code blocks.

---

## 1. System Models and Foundations

- **[Indistinguishability of Crash vs Delay](../00-foundations/system-models.md#4-the-indistinguishability-principle)**: Sequence diagram showing how network delay and crash failures produce identical local observations.
- **[Global Stabilization Time (GST) Timeline](../00-foundations/synchrony-and-asynchrony.md#51-formal-formulations)**: Visualizing the transition from asynchronous behavior to bounded delay.
- **[Byzantine Generals Impossibility ($n=3, f=1$)](../00-foundations/failure-models.md#32-byzantine-fault-threshold-n-ge-3f--1-or-f--n3)**: Sequence diagram demonstrating indistinguishability between a traitorous commander and a traitorous lieutenant.
- **[Two Generals Regress](../00-foundations/impossibility-results.md#11-the-model-and-problem-formulation)**: Sequence trace showing why acknowledgments of acknowledgments cannot establish common knowledge.
- **[The CAP Partition Barrier](../00-foundations/impossibility-results.md#32-the-proof-by-indistinguishability)**: Partition trace forcing a choice between Linearizability and Availability.

---

## 2. Time, Ordering, and Coordination

- **[Physical Clock Drift Curve](../01-time-and-ordering/physical-clocks.md#21-quartz-crystal-oscillators)**: XY chart showing linear accumulation of clock skew.
- **[NTP Round-Trip Skew Estimation](../01-time-and-ordering/physical-clocks.md#5-network-time-protocol-ntp)**: Sequence diagram showing $t_1, t_2, t_3, t_4$ timestamps and transit delay midpoint.
- **[Google Spanner TrueTime Commit Wait](../01-time-and-ordering/physical-clocks.md#62-the-commit-wait-rule)**: Sequence diagram proving external consistency via uncertainty interval pause.
- **[Happens-Before Space-Time Diagram](../01-time-and-ordering/happens-before.md#5-worked-example-and-space-time-diagram)**: Visualizing local program order, message causality, and concurrent events.
- **[Relativistic Light Cones in Distributed Systems](../01-time-and-ordering/happens-before.md#6-causal-pasts-futures-and-relativistic-light-cones)**: Graphing past, future, and concurrent causal regions.
- **[Vector Clocks Causality Trace](../01-time-and-ordering/vector-clocks.md#5-worked-example-step-by-step-execution-trace)**: Sequence diagram tracing 3-process vector evolution and concurrency detection.
- **[Quorum Intersection & Read/Write Overlap](../02-communication-and-coordination/quorums.md#31-the-read-write-quorum-invariant)**: Visualizing $r + w > n$ pigeonhole intersection.
- **[3x3 Grid Quorum Topology](../02-communication-and-coordination/quorums.md#51-grid-matrix-quorums)**: Visualizing row-column intersection for $O(\sqrt{n})$ messages.
- **[Quorum Staleness Inversion](../02-communication-and-coordination/quorums.md#7-the-quorum-staleness-hazard)**: Demonstrating how naive quorum reads can return stale values.
- **[Failure Detector Accrual & Adaptation](../02-communication-and-coordination/failure-detectors.md#5-practical-implementation-heartbeats-and-adaptive-timeouts)**: Sequence trace of heartbeat timeouts and recovery.
- **[Fencing Tokens Guarding Against Stale Leaders](../02-communication-and-coordination/leader-election.md#5-fencing-tokens-guarding-against-zombie-leaders)**: Sequence trace illustrating how epoch tokens block deposed leaders.

---

## 3. Replication, Consensus, and Transactions

- **[Lamport Register Overlap Intervals](../03-replication-and-consistency/shared-registers.md#21-operation-intervals-and-precedence)**: Timeline showing concurrent vs non-overlapping read/write intervals.
- **[Regular Register Flicker Anomaly](../03-replication-and-consistency/shared-registers.md#32-regular-register)**: Sequence diagram showing a newer read followed by an older read.
- **[The ABD Read Write-Back Sequence](../03-replication-and-consistency/abd-algorithm.md#42-protocol-phases)**: Two-phase query and write-back sequence diagram.
- **[Stale Read Inversion without Write-Back](../03-replication-and-consistency/abd-algorithm.md#5-why-read-write-back-is-strictly-necessary)**: Step-by-step failure sequence proving why Phase 2 is strictly necessary.
- **[Linearization Point Mapping](../03-replication-and-consistency/linearizability.md#3-the-linearization-point)**: Sequence trace showing real-time bounds and linearization points.
- **[Non-Linearizable Execution Gantt Chart](../03-replication-and-consistency/linearizability.md#61-a-non-linearizable-execution-trace-the-stale-inversion-anomaly)**: Gantt chart proving contradiction in sequential order.
- **[State Machine Replication Flow](../03-replication-and-consistency/state-machine-replication.md#3-the-core-architecture-smr-via-total-order-broadcast)**: Sequence diagram from client to total order broadcast to state machines.
- **[Paxos Phase 1 & Phase 2 Sequence](../04-consensus/paxos.md#4-single-decree-paxos-protocol-specification)**: Complete Synod message flow (Prepare/Promise, Accept/Accepted).
- **[Dueling Proposers Livelock](../04-consensus/paxos.md#6-livelock-hazard-dueling-proposers)**: Sequence diagram of alternating ballot preemption.
- **[Raft Server Role State Machine](../04-consensus/raft.md#2-server-roles-and-state-transitions)**: State diagram of Follower, Candidate, and Leader transitions.
- **[Raft Figure 8 Commit Hazard](../04-consensus/raft.md#6-the-subtle-commit-rule-committing-entries-from-earlier-terms)**: Gantt chart showing why leaders cannot commit older terms by counting replicas.
- **[PBFT Three-Phase Commit Sequence](../04-consensus/byzantine-fault-tolerance.md#4-the-three-phase-pbft-protocol)**: Pre-Prepare, Prepare, and Commit message flow with $3f+1$ nodes.
- **[Classical Two-Phase Commit (2PC)](../05-storage-and-transactions/two-phase-commit.md#3-the-classical-two-phase-commit-protocol-gray-1978)**: Sequence diagram of prepare and commit phases.
- **[Distributed Deadlock Cycle](../05-storage-and-transactions/distributed-transactions.md#3-distributed-two-phase-locking-2pl-and-deadlocks)**: Sequence diagram showing cross-shard circular lock waiting.
- **[Consistent Hashing Virtual Ring](../05-storage-and-transactions/consistent-hashing.md#42-the-virtual-node-solution-dynamo-2007)**: Interleaved virtual token ring topology.
- **[Chandy-Lamport Distributed Snapshot](../05-storage-and-transactions/recovery-and-rebalancing.md#21-consistent-global-checkpointing-chandy-lamport-algorithm-1985)**: Marker propagation and channel recording.

---

## 4. Production Engineering

- **[Circuit Breaker State Machine](../06-production-systems/timeouts-and-retries.md#4-the-circuit-breaker-pattern-nygard-2007)**: Closed, Open, and Half-Open lifecycle.
- **[Idempotency Gateway Sequence](../06-production-systems/idempotency.md#3-idempotency-keys-deduplication-tokens)**: Sequence diagram showing duplicate request deduplication and response caching.
- **[Token Bucket Flow Control](../06-production-systems/backpressure.md#31-the-token-bucket-algorithm)**: Rate limiting token refill and consumption flow.
- **[The Cascading Failure Death Spiral](../06-production-systems/load-balancing.md#42-the-cascading-failure-death-spiral)**: Flowchart of load balancer failover overload.
- **[Distributed Tracing Span DAG](../06-production-systems/observability.md#2-distributed-tracing-and-context-propagation)**: Gantt chart of child and grandchild span causality.
- **[Tail Latency Amplification Curve](../06-production-systems/observability.md#3-the-mathematics-of-tail-latency-dean--barroso-2013)**: Chart showing user delay probability rising to $63\%$ across 100 shards.
