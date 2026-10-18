# Authoritative References and Bibliography

This bibliography compiles seminal papers, authoritative textbooks, and foundational monographs underpinning modern distributed systems theory. Every note in this repository derives its formal models, definitions, and proofs from these sources.

---

## 1. Primary Textbooks and Monograph Collections

1. **[Cachin, Guerraoui, Rodrigues 2011]**  
   Christian Cachin, Rachid Guerraoui, and Luís Rodrigues.  
   *Introduction to Reliable and Secure Distributed Programming* (2nd Edition). Springer, 2011.  
   *Significance*: The definitive mathematical treatise on modular distributed algorithms, failure detectors, reliable broadcast abstractions, shared registers (ABD), and consensus in crash and Byzantine models.

2. **[Kleppmann 2017]**  
   Martin Kleppmann.  
   *Designing Data-Intensive Applications: The Big Ideas Behind Reliable, Scalable, and Maintainable Systems*. O'Reilly Media, 2017.  
   *Significance*: Bridges theoretical fundamentals (linearizability, total order broadcast, consensus) with production storage architectures (LSM-trees, B-trees, replication topologies, transactions).

3. **[Kleppmann 2021]**  
   Martin Kleppmann.  
   *Distributed Systems Lecture Notes* (Computer Science Tripos, Part IB). University of Cambridge, 2021/22.  
   *Significance*: Foundational pedagogical treatment of models, physical/logical time, broadcast hierarchies, quorums, Raft, 2PC, CRDTs, and Spanner TrueTime.

4. **[Lynch 1996]**  
   Nancy A. Lynch.  
   *Distributed Algorithms*. Morgan Kaufmann Publishers, 1996.  
   *Significance*: The canonical formal foundation using I/O Automata to prove correctness and complexity bounds for synchronous, asynchronous, and shared-memory algorithms.

5. **[van Steen, Tanenbaum 2017]**  
   Maarten van Steen and Andrew S. Tanenbaum.  
   *Distributed Systems* (3rd Edition). CreateSpace Independent Publishing, 2017.  
   *Significance*: Comprehensive treatment of architectural patterns, communication protocols, naming, synchronization, and fault tolerance.

6. **[Attiya, Welch 2004]**  
   Hagit Attiya and Jennifer Welch.  
   *Distributed Computing: Fundamentals, Simulations, and Advanced Topics* (2nd Edition). Wiley-Interscience, 2004.  
   *Significance*: Exhaustive mathematical coverage of shared-memory register hierarchies, consensus bounds, and synchronization primitives.

---

## 2. Seminal Research Papers

### Foundations, Time, and Ordering

- **[Lamport 1978]** Leslie Lamport.  
  *Time, Clocks, and the Ordering of Events in a Distributed System*.  
  Communications of the ACM (CACM), 21(7):558–565, July 1978.  
  *Introduced*: The happens-before relation ($\to$), logical clocks, and distributed mutual exclusion via total ordering.

- **[Mattern 1989]** Friedemann Mattern.  
  *Virtual Time and Global States of Distributed Systems*.  
  Workshop on Parallel and Distributed Algorithms, pp. 215–226, 1989.  
  *Introduced*: Vector time, causal cuts, and lattice models of distributed executions.

- **[Fidge 1988]** Colin J. Fidge.  
  *Timestamps in Message-Passing Systems That Preserve the Partial Ordering*.  
  Australian Computer Science Communications, 10(1):56–66, 1988.  
  *Concurrent discovery*: Dynamic vector timestamping for concurrent process tracking.

- **[Kulkarni et al. 2014]** Sandeep S. Kulkarni, Murat Demirbas, Deepak Madeppa, Bharadwaj Avva, and Marcelo Leone.  
  *Logical Physical Clocks and Consistent Snapshots in Globally Distributed Databases*.  
  State University of New York at Buffalo Technical Report, 2014 / OPODIS 2014.  
  *Introduced*: Hybrid Logical Clocks (HLC) combining physical NTP tracking with Lamport causality bounds.

### Impossibility Results

- **[Gray 1978]** Jim N. Gray.  
  *Notes on Data Base Operating Systems*.  
  Operating Systems, LNCS 60, pp. 393–481. Springer-Verlag, 1978.  
  *Introduced*: The Two Generals Problem, atomic commit, and Two-Phase Commit (2PC).

- **[Fischer, Lynch, Paterson 1985] (FLP)**  
  Michael J. Fischer, Nancy A. Lynch, and Michael S. Paterson.  
  *Impossibility of Distributed Consensus with One Faulty Process*.  
  Journal of the ACM (JACM), 32(2):374–382, April 1985.  
  *Core Result*: Proves that in an asynchronous message-passing system, no deterministic consensus protocol can guarantee liveness in the presence of even a single unannounced crash failure.

- **[Dwork, Lynch, Stockmeyer 1988] (DLS)**  
  Cynthia Dwork, Nancy A. Lynch, and Larry Stockmeyer.  
  *Consensus in the Presence of Partial Synchrony*.  
  Journal of the ACM (JACM), 35(2):288–323, April 1988.  
  *Introduced*: Formal models of partial synchrony ($\Delta$ known after Global Stabilization Time GST, or bounds exist but are unknown) and consensus feasibility.

- **[Gilbert, Lynch 2002]**  
  Seth Gilbert and Nancy A. Lynch.  
  *Brewer's Conjecture and the Feasibility of Consistent, Available, Partition-Tolerant Web Services*.  
  ACM SIGACT News, 33(2):51–59, June 2002.  
  *Core Result*: Formal proof that an asynchronous read/write system cannot simultaneously provide linearizability ($C$) and availability ($A$) under network partitions ($P$).

- **[Kleppmann 2015]**  
  Martin Kleppmann.  
  *A Critique of the CAP Theorem*.  
  arXiv:1509.05393, September 2015.  
  *Clarification*: Nuanced deconstruction of 1-safe vs 2-safe register models, network partitions vs latency spikes, and PACELC trade-offs.

### Coordination, Failure Detectors, and Quorums

- **[Chandra, Toueg 1996]**  
  Tushar Deepak Chandra and Sam Toueg.  
  *Unreliable Failure Detectors for Reliable Distributed Systems*.  
  Journal of the ACM (JACM), 43(2):225–267, March 1996.  
  *Introduced*: Failure detector classes ($\mathcal{P}, \mathcal{S}, \diamondsuit\mathcal{P}, \diamondsuit\mathcal{S}, \Omega$), completeness and accuracy metrics, and reduction to consensus.

- **[Gifford 1979]** David K. Gifford.  
  *Weighted Voting for Replicated Data*.  
  Proceedings of the 7th ACM Symposium on Operating Systems Principles (SOSP), pp. 150–162, 1979.  
  *Introduced*: Quorum consensus with read/write quorum configurations ($r + w > n$).

- **[Whittaker et al. 2021]**  
  Michael Whittaker, Aleksey Charapko, Joseph M. Hellerstein, Heidi Howard, and Ion Stoica.  
  *Read-Write Quorum Systems Made Practical*.  
  PaPoC 2021.  
  *Introduced*: Dynamic, asymmetric, and heterogeneous grid quorum constructions beyond simple majorities.

### Shared Registers and Linearizability

- **[Lamport 1986]** Leslie Lamport.  
  *On Interprocess Communication. Part I: Basic Formalisms; Part II: Algorithms*.  
  Distributed Computing, 1(2):77–101, 1986.  
  *Introduced*: Formal definitions of Safe, Regular, and Atomic registers for concurrent read/write executions.

- **[Herlihy, Wing 1990]**  
  Maurice P. Herlihy and Jeannette M. Wing.  
  *Linearizability: A Correctness Condition for Concurrent Objects*.  
  ACM Transactions on Programming Languages and Systems (TOPLAS), 12(3):463–492, July 1990.  
  *Introduced*: Linearizability, real-time precedence ordering ($op_1 \prec_H op_2$), sequential specifications, and local property (composability) proof.

- **[Attiya, Bar-Noy, Dolev 1995] (ABD)**  
  Hagit Attiya, Amotz Bar-Noy, and Danny Dolev.  
  *Sharing Memory Robustly in Message-Passing Systems*.  
  Journal of the ACM (JACM), 42(1):124–142, January 1995.  
  *Introduced*: The ABD algorithm implementing linearizable SWMR and MWMR atomic registers over asynchronous message-passing quorum systems with crash-stop failures ($f < n/2$).

- **[Lynch, Shvartsman 1997] (RAMBO)**  
  Nancy A. Lynch and Alex A. Shvartsman.  
  *Robust Emulation of Shared Memory Using Dynamic Quorum-Acknowledged Broadcasts*.  
  FTCS, pp. 272–281, 1997.  
  *Extension*: Emulating atomic registers across dynamically changing node configurations.

### Consensus Protocols: Paxos, Raft, and Byzantine Fault Tolerance

- **[Lamport 1998]** Leslie Lamport.  
  *The Part-Time Parliament*.  
  ACM Transactions on Computer Systems (TOCS), 16(2):133–169, May 1998 (submitted 1990).  
  *Introduced*: The Synod protocol and Multi-Paxos consensus.

- **[Lamport 2001]** Leslie Lamport.  
  *Paxos Made Simple*.  
  ACM SIGACT News, 32(4):18–25, December 2001.  
  *Pedagogical*: Direct constructive derivation of Phase 1 (Prepare/Promise) and Phase 2 (Accept/Accepted).

- **[Ongaro, Ousterhout 2014]**  
  Diego Ongaro and John Ousterhout.  
  *In Search of an Understandable Consensus Algorithm*.  
  USENIX Annual Technical Conference (ATC), pp. 305–319, June 2014.  
  *Introduced*: The Raft consensus protocol, log matching invariants, leader completeness, and joint consensus cluster membership changes.

- **[Castro, Liskov 2002] (PBFT)**  
  Miguel Castro and Barbara Liskov.  
  *Practical Byzantine Fault Tolerance and Proactive Recovery*.  
  ACM Transactions on Computer Systems (TOCS), 20(4):398–461, November 2002.  
  *Introduced*: The first practical state-machine replication algorithm tolerating Byzantine faults in asynchronous networks with $3f+1$ nodes, using pre-prepare, prepare, and commit phases.

- **[Howard, Mortier 2020]**  
  Heidi Howard and Richard Mortier.  
  *Paxos vs Raft: Have we reached consensus on distributed consensus?*  
  PaPoC 2020.  
  *Comparative*: In-depth theoretical dissection of Paxos and Raft mechanisms, leader elections, and quorum utilization.

### Distributed Storage, Transactions, and Consistency Models

- **[Corbett et al. 2012]**  
  James C. Corbett, Jeffrey Dean, Michael Epstein, Andrew Fikes, Christopher Frost, et al.  
  *Spanner: Google’s Globally-Distributed Database*.  
  10th USENIX OSDI, pp. 251–264, October 2012.  
  *Introduced*: TrueTime API ($[t_{earliest}, t_{latest}]$ uncertainty bounds using GPS and atomic clocks), commit wait, external consistency (strict serializability), and lock-free read-only snapshot transactions.

- **[DeCandia et al. 2007]**  
  Giuseppe DeCandia, Deniz Hastorun, Madan Jampani, Gunavardhan Kakulapati, et al.  
  *Dynamo: Amazon’s Highly Available Key-value Store*.  
  21st ACM SOSP, pp. 205–220, October 2007.  
  *Introduced*: Consistent hashing with virtual nodes, sloppy quorums, hinted handoff, vector clocks for version reconciliation, and anti-entropy via Merkle trees.

- **[Shapiro et al. 2011]**  
  Marc Shapiro, Nuno Preguiça, Carlos Baquero, and Marek Zawirski.  
  *Conflict-Free Replicated Data Types*.  
  SSS 2011, LNCS 6976, pp. 386–400. Springer, 2011.  
  *Introduced*: Formal definition of Strong Eventual Consistency (SEC), CvRDT (state-based monotonic join semilattices), and CmRDT (operation-based commutative protocols).

- **[Gray, Lamport 2006]**  
  Jim Gray and Leslie Lamport.  
  *Consensus on Transaction Commit*.  
  ACM Transactions on Database Systems (TODS), 31(1):133–160, March 2006.  
  *Introduced*: Paxos Commit, eliminating the coordinator single point of failure in 2PC by replacing the coordinator with a Paxos consensus group.
