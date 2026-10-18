# Distributed Systems Theory Glossary

A rigorous reference glossary of mathematical definitions, formal invariants, and operational abstractions across distributed computing theory.

---

## 1. System Models, Processes, and Executions

### Distributed System
A collection of autonomous computing entities (processes/nodes) $\Pi = \{p_1, p_2, \dots, p_n\}$ that do not share physical memory or a centralized physical clock, communicating exclusively by exchanging messages over an interconnection network.

### Execution History ($\mathcal{H}$)
A sequence of events occurring across processes. An event $e \in \mathcal{E}$ represents an atomic state transition at a process: an internal execution step, a message send event $\text{send}(m)$, or a message invocation/receipt event $\text{recv}(m)$.

### Synchrony Models
- **Synchronous System ($\mathcal{SYNCH}$)**: A system where there exists a known, fixed upper bound $\Delta$ on message transmission delay, and known bounds $[\Phi_{\min}, \Phi_{\max}]$ on process execution step times.
- **Asynchronous System ($\mathcal{ASYNCH}$)**: A system where message transmission latency and process execution speeds are finite but arbitrarily large and unpredictable. No assumptions are made regarding physical clock synchronization or transmission bounds.
- **Partially Synchronous System ($\mathcal{PARTIAL}$)**: A system where timing bounds exist but are either unknown *a priori*, or hold only after an unknown but finite point in time known as the **Global Stabilization Time (GST)** [Dwork, Lynch, Stockmeyer 1988].

### Process Failure Models
- **Crash-Stop (Fail-Stop)**: A process executes faithfully according to its protocol specification until it potentially crashes. Once crashed, it halts all execution permanently and never sends or receives messages again.
- **Crash-Recovery (Fail-Recovery)**: A process may crash at any arbitrary point, losing volatile memory state, but may subsequently restart. Persistent state stored on non-volatile storage (disk/WAL) survives crashes.
- **Byzantine (Fail-Arbitrary)**: A faulty process may deviate from its prescribed algorithm in arbitrary, non-deterministic, or coordinated adversarial ways (e.g., sending contradictory messages, colluding with other faulty nodes, forging metadata). A non-faulty process is termed **correct** or **honest**.

### Network Link Models
- **Reliable (Perfect) Links**: A link satisfying:
  1. *Validity (No loss)*: If a correct process $p$ sends message $m$ to correct process $q$, then $q$ eventually receives $m$.
  2. *Integrity (No creation/duplication)*: A message $m$ is received by $q$ at most once, and only if $m$ was previously sent by $p$.
- **Fair-Loss Links**: Messages may be dropped, duplicated, or reordered in flight. However, if a message $m$ is retransmitted infinitely often between two correct processes, $m$ is received infinitely often.
- **Arbitrary (Adversarial) Links**: An active network adversary may drop, delay, duplicate, tamper with, or inject spurious packets into communication channels.

---

## 2. Properties: Safety, Liveness, and Correctness

### Property ($\mathcal{P}$)
A set of infinite execution traces over system events. An execution satisfies property $\mathcal{P}$ if the execution trace is an element of $\mathcal{P}$.

### Safety Property
A property stating that "something bad never happens" [Alpern and Schneider 1985]. Formally, $\mathcal{P}$ is a safety property if for every execution trace $\sigma \notin \mathcal{P}$, there exists a finite prefix $\sigma_{pref}$ of $\sigma$ such that no infinite extension $\sigma_{pref} \cdot \beta$ can ever satisfy $\mathcal{P}$. Irremediable once violated.

### Liveness Property
A property stating that "something good eventually happens" [Alpern and Schneider 1985]. Formally, $\mathcal{P}$ is a liveness property if every finite execution prefix $\sigma_{pref}$ can be extended by some infinite sequence $\beta$ such that $\sigma_{pref} \cdot \beta \in \mathcal{P}$. No finite trace can definitively prove a violation of liveness.

### Invariant
A predicate over the global system state that evaluates to true in every reachable state from the initial configuration across all valid execution schedules.

---

## 3. Time, Ordering, and Causality

### Happens-Before Relation ($\to$)
The strict partial order on events defined by Leslie Lamport [1978]:
1. If $a$ and $b$ are events in the same process, and $a$ occurs before $b$, then $a \to b$.
2. If event $a$ is the sending of a message $m$ by process $p_i$, and event $b$ is the receipt of $m$ by process $p_j$, then $a \to b$.
3. If $a \to c$ and $c \to b$, then $a \to b$ (transitivity).

### Concurrency ($a \parallel b$)
Two events $a$ and $b$ are concurrent ($a \parallel b$) if and only if neither causally precedes the other:
$$a \parallel b \iff (a \not\to b) \land (b \not\to a)$$

### Lamport Logical Clock
A monotonically increasing scalar function $L: \mathcal{E} \to \mathbb{N}$ ensuring the **Clock Condition**:
$$a \to b \implies L(a) < L(b)$$
*Limitation*: The converse does not hold ($L(a) < L(b) \not\implies a \to b$).

### Vector Clock
An array $V \in \mathbb{N}^n$ tracking causal histories across $n$ processes:
$$V(a) \le V(b) \iff \forall k \in \{1,\dots,n\}: V(a)[k] \le V(b)[k]$$
$$a \to b \iff (V(a) \le V(b)) \land (V(a) \ne V(b))$$
Vector clocks provide an exact isomorphism to the happens-before partial order.

### Clock Drift and Clock Skew
- **Drift Rate ($\rho$)**: The rate at which an uncalibrated physical quartz oscillator deviates from true International Atomic Time (TAI): $\left| \frac{dC}{dt} - 1 \right| \le \rho$ (typically $\approx 10^{-5}$ to $10^{-6}$ or $10\text{--}50\text{ ppm}$).
- **Clock Skew ($\theta$)**: The instantaneous physical time difference between two clocks $C_i(t)$ and $C_j(t)$ at a given instant $t$: $\theta = |C_i(t) - C_j(t)|$.
- **Slewing**: Adjusting clock frequency gradually without causing discrete backwards jumps.
- **Stepping**: Forcibly updating clock time discretely; can cause non-monotonicity.

---

## 4. Coordination, Broadcast, and Quorums

### Reliable Broadcast (FIFO, Causal, Total Order)
- **FIFO Broadcast**: If process $p$ broadcasts $m_1$ before $m_2$, all correct processes deliver $m_1$ before $m_2$.
- **Causal Broadcast**: If $\text{broadcast}(m_1) \to \text{broadcast}(m_2)$, all correct processes deliver $m_1$ before $m_2$.
- **Total Order Broadcast (Atomic Broadcast)**: If any process delivers $m_1$ before $m_2$, every process delivers $m_1$ before $m_2$. Formally equivalent to consensus [Chandra, Toueg 1996].

### Quorum System
A collection $\mathcal{Q} = \{Q_1, Q_2, \dots, Q_m\}$ of subsets of processes $\Pi$ such that every pair of quorums intersects:
$$\forall Q_j, Q_k \in \mathcal{Q}: Q_j \cap Q_k \ne \emptyset$$
In a simple majority quorum system over $n$ processes, each quorum has size $\lfloor n/2 \rfloor + 1$. For read-write quorums with write quorums of size $w$ and read quorums of size $r$, the intersection invariant requires:
$$r + w > n$$

### Failure Detector ($\mathcal{D}$)
A distributed oracle that provides processes with hints about which nodes are suspected of having crashed:
- **Completeness**: Strong completeness requires that eventually every crashed process is suspected by every correct process.
- **Accuracy**: Strong accuracy requires that no correct process is ever suspected. Eventual strong accuracy ($\diamondsuit\mathcal{P}$) allows spurious suspicions during unstable periods, but eventually no correct process is suspected.

---

## 5. Replication and Consistency Models

### Shared Register Hierarchy [Lamport 1986]
- **Safe Register**: A read that does not overlap with any write returns the most recently written value. A read that overlaps with a write may return any arbitrary value within the register domain.
- **Regular Register**: A read returns either the value written by the most recent non-overlapping write, or the value of any write overlapping with the read.
- **Atomic Register (Linearizable Register)**: Operations take effect instantaneously at a discrete linearization point between their invocation and response. Every read returns the value of the most recent write in the resulting total order.

### Linearizability [Herlihy, Wing 1990]
A consistency condition where each operation appears to take effect instantaneously at some point $t_{linear}$ between its invocation event $\text{inv}(op)$ and response event $\text{resp}(op)$ in global real time. Linearizability is a **local property** (composable across disjoint objects).

### Sequential Consistency [Lamport 1979]
The result of any execution is the same as if the operations of all processors were executed in some sequential order, and the operations of each individual processor appear in this sequence in the order specified by its program. Unlike linearizability, sequential consistency is not constrained by real-time precedence and is **not composable**.

### Eventual Consistency & Strong Eventual Consistency (SEC)
- **Eventual Consistency**: If no new updates are submitted, all replicas eventually converge to the same state.
- **Strong Eventual Consistency (SEC)** [Shapiro et al. 2011]: Guarantees that any two replicas that have delivered the same set of updates are in identical state, without requiring rollback or consensus. Implemented by Conflict-Free Replicated Data Types (CRDTs).

### State Machine Replication (SMR) [Lamport 1984, Schneider 1990]
A general paradigm where multiple replicas maintain identical state machines starting from identical initial states, transitioning deterministically on a sequence of inputs delivered via Total Order Broadcast.

---

## 6. Consensus and Impossibility Theorems

### Consensus Problem
A protocol where each process $p_i$ proposes an initial value $v_i \in \mathcal{V}$ and must decide on a value $d \in \mathcal{V}$ satisfying:
1. **Agreement**: No two correct processes decide different values.
2. **Validity (Integrity)**: If a correct process decides $d$, then $d$ must have been proposed by some process.
3. **Termination**: Every correct process eventually decides a value.

### FLP Impossibility Result [Fischer, Lynch, Paterson 1985]
In an asynchronous message-passing system ($\mathcal{ASYNCH}$), no deterministic algorithm can solve consensus if even a single process is subject to unannounced crash-stop failure ($f \ge 1$).

### CAP Theorem [Fox-Brewer 1999, Gilbert-Lynch 2002]
In an asynchronous read/write data store, it is impossible to simultaneously guarantee Linearizability (Consistency, $C$) and 100% Availability ($A$) in the presence of an arbitrary Network Partition ($P$).

---

## 7. Distributed Transactions and Storage

### Two-Phase Commit (2PC) [Gray 1978]
An atomic commitment protocol ensuring that all participants either commit or abort a distributed transaction:
- **Phase 1 (Prepare)**: Coordinator queries participants whether they can commit. Participants acquire locks, write undo/redo logs, and vote `YES` or `NO`.
- **Phase 2 (Commit/Abort)**: If all vote `YES`, coordinator commits; otherwise, coordinator aborts.
- *Limitation*: The coordinator is a single point of failure. If the coordinator crashes while participants are in the prepared state, participants remain **in-doubt** (blocked).

### Multi-Version Concurrency Control (MVCC)
A concurrency control method where write operations create a new immutable, timestamped version of a record rather than overwriting existing data in place. Read operations access a historical snapshot identified by a timestamp $T_{read}$, ensuring non-blocking reads without read-locks.

### External Consistency (Strict Serializability)
The combination of Serializability (transactions behave as if executed in some serial order) and Linearizability (if transaction $T_1$ commits in real time before transaction $T_2$ begins, $T_1$ must precede $T_2$ in the serial order).
