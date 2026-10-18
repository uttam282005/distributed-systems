# Repository Development Progress

This document tracks the authoring, technical auditing, and completion status of all topics in the **Distributed Systems Theory Notes Repository**.

---

## 1. Status Legend

- 🟢 **Complete (Substantive)**: Fully documented with formal definitions, algorithm pseudocode, execution traces, proofs/invariants, failure modes, complexity bounds, and authoritative references.
- 🟡 **In Progress**: Drafted or partially completed; currently being expanded.
- ⚪ **Planned**: Outlined and scheduled in dependency sequence.

---

## 2. Topic Audit Tracker

| Category | Note / File | Status | Key Topics Covered |
| :--- | :--- | :---: | :--- |
| **Meta** | `README.md` | 🟢 | Repo architecture, pedagogical standards, navigation |
| **Meta** | `ROADMAP.md` | 🟢 | 7-Stage dependency-aware learning sequence & self-check criteria |
| **Meta** | `GLOSSARY.md` | 🟢 | Formal mathematical definitions across all modules |
| **Meta** | `references/README.md` | 🟢 | Authoritative primary sources, textbooks, and seminal papers |
| **00 Foundations** | `distributed-systems-overview.md` | 🟢 | Core definition, concurrency, partial failure, physical vs distributed |
| **00 Foundations** | `system-models.md` | 🟢 | Processes, execution histories, network links, adversaries |
| **00 Foundations** | `synchrony-and-asynchrony.md` | 🟢 | $\mathcal{SYNCH}$, $\mathcal{ASYNCH}$, $\mathcal{PARTIAL}$, GST, practical violations |
| **00 Foundations** | `failure-models.md` | 🟢 | Crash-stop, crash-recovery, Byzantine, correctness thresholds |
| **00 Foundations** | `safety-and-liveness.md` | 🟢 | Alpern-Schneider trace characterization, finite prefix refutations |
| **00 Foundations** | `impossibility-results.md` | 🟢 | Two Generals problem, FLP formulation, CAP, lower bounds |
| **01 Time & Ordering** | `physical-clocks.md` | 🟢 | Quartz drift, NTP Cristian/Marzullo skew estimation, TrueTime |
| **01 Time & Ordering** | `happens-before.md` | 🟢 | Lamport's $\to$, strict partial order proof, concurrency $\parallel$ |
| **01 Time & Ordering** | `lamport-clocks.md` | 🟢 | Scalar logical clocks, Clock Condition, tie-breaking total orders |
| **01 Time & Ordering** | `vector-clocks.md` | 🟢 | Vector clocks, causal history isomorphism, causality verification |
| **02 Communication** | `message-passing.md` | 🟢 | Unicast messaging, FIFO channel construction, deduplication |
| **02 Communication** | `rpc.md` | 🟢 | Remote Procedure Calls, stubs, failure semantics, location transparency |
| **02 Communication** | `reliable-broadcast.md` | 🟢 | Best-effort, Reliable, FIFO, Causal, and Total Order Broadcast |
| **02 Communication** | `failure-detectors.md` | 🟢 | Chandra-Toueg classes ($\mathcal{P}, \diamondsuit\mathcal{P}, \diamondsuit\mathcal{S}, \Omega$), completeness, accuracy |
| **02 Communication** | `leader-election.md` | 🟢 | Symmetry breaking, Bully algorithm, ring election, split-brain |
| **02 Communication** | `quorums.md` | 🟢 | Quorum systems, majority quorums, $r+w > n$, Byzantine quorums |
| **03 Replication** | `replication-fundamentals.md` | 🟢 | Active/passive replication, state vs operation transfer, anti-entropy |
| **03 Replication** | `primary-backup.md` | 🟢 | Leader-follower replication, failover hazards, split-brain |
| **03 Replication** | `shared-registers.md` | 🟢 | Lamport register hierarchy: Safe, Regular, and Atomic registers |
| **03 Replication** | `abd-algorithm.md` | 🟢 | SWMR and MWMR ABD, query & write-back phases, stale read proof |
| **03 Replication** | `linearizability.md` | 🟢 | Herlihy-Wing linearizability, linearization points, composability |
| **03 Replication** | `sequential-consistency.md` | 🟢 | Sequential consistency, program order, non-composability |
| **03 Replication** | `eventual-consistency.md` | 🟢 | Optimistic replication, LWW, multi-value registers, CRDTs |
| **03 Replication** | `state-machine-replication.md` | 🟢 | Deterministic state machines, total order broadcast equivalence |
| **04 Consensus** | `consensus-problem.md` | 🟢 | Agreement, Validity, Termination, relationship to Atomic Broadcast |
| **04 Consensus** | `flp-impossibility.md` | 🟢 | Fischer-Lynch-Paterson proof, bivalent configurations, commutativity |
| **04 Consensus** | `paxos.md` | 🟢 | Single-decree Synod, Phase 1 & 2, Multi-Paxos |
| **04 Consensus** | `raft.md` | 🟢 | Leader election, log replication, commit rule, safety invariants |
| **04 Consensus** | `byzantine-fault-tolerance.md` | 🟢 | Byzantine Generals, PBFT 3-phase flow, $3f+1$ bound |
| **05 Storage & Txns** | `partitioning-and-sharding.md` | 🟢 | Range vs hash partitioning, secondary indices, routing |
| **05 Storage & Txns** | `consistent-hashing.md` | 🟢 | Hash ring topology, virtual nodes, $O(K/N)$ rebalancing |
| **05 Storage & Txns** | `replicated-logs.md` | 🟢 | Write-Ahead Logging, LSM-trees, segment compaction |
| **05 Storage & Txns** | `two-phase-commit.md` | 🟢 | 2PC protocol, coordinator failure, in-doubt states, Paxos Commit |
| **05 Storage & Txns** | `distributed-transactions.md` | 🟢 | Distributed ACID, 2PL, optimistic concurrency control, SSI |
| **05 Storage & Txns** | `mvcc.md` | 🟢 | Multi-version concurrency control, snapshot isolation, TrueTime |
| **05 Storage & Txns** | `recovery-and-rebalancing.md` | 🟢 | Checkpointing, crash recovery, node addition/removal |
| **06 Production** | `timeouts-and-retries.md` | 🟢 | Exponential backoff, full jitter, deadline propagation |
| **06 Production** | `idempotency.md` | 🟢 | Idempotency ($f(f(x))=f(x)$), deduplication tokens, tombstones |
| **06 Production** | `backpressure.md` | 🟢 | Flow control, reactive drop, sliding windows, token bucket |
| **06 Production** | `load-balancing.md` | 🟢 | L4/L7 routing, power-of-two-choices, health check herd effects |
| **06 Production** | `service-discovery.md` | 🟢 | Dynamic registries, gossip discovery, heartbeat leases |
| **06 Production** | `observability.md` | 🟢 | Distributed tracing, W3C Trace Context, causal propagation |
| **06 Production** | `disaster-recovery.md` | 🟢 | Multi-region replication, RPO/RTO invariants, cross-datacenter failover |

---

## 3. Milestone Summary

- **Phase 1: Architecture & Foundations**: Structure, central navigation, roadmap, formal glossary, bibliography.
- **Phase 2: Core Theoretical Treatises**: Deep mathematical specifications of quorums, linearizability, and the ABD algorithm (both SWMR & MWMR), consensus, and Raft.
- **Phase 3: Comprehensive Topic Rollout**: Complete coverage of all 42 modules across Foundations, Time, Coordination, Replication, Consensus, Storage, and Production Systems.
- **Phase 4: Audit & Verification**: Internal cross-link validation, Mermaid syntax check, and git commit history tracking.
