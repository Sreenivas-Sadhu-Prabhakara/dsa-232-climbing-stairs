# Climbing Stairs

> NeetCode 150 · 1-D DP · Blind 75 — the gateway DP: the last move was 1 or 2 steps, so ways(n) = ways(n−1) + ways(n−2), i.e. Fibonacci in two variables. Go, stdlib only, tested.

## Problem

Climb n stairs taking 1 or 2 steps at a time; count the distinct step sequences. The canonical first dynamic-programming problem.

## Approach

Condition on the final move: every route to step n arrives from n−1 (last hop 1) or n−2 (last hop 2), and those route sets are disjoint — so the counts add. That recurrence is Fibonacci with shifted seeds (ways(0)=ways(1)=1). Each state needs only its two predecessors, so two rolling variables replace the memo array — the space-collapse that most 1-D DPs on this list repeat (house robber, decode ways). Naive recursion without memoization is the exponential trap the problem exists to teach out of you.

**Complexity:** O(n) time, O(1) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
