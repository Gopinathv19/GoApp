---
gitfill: "brainstorm"
title: "Group booking in one go"
source_commit: "cfc0167"
generated: "2026-11-02"
status: "ready"
---

# Group booking in one go

> Implementation spec from a GitFill brainstorm, grounded in `Gopinathv19/GoApp` at `cfc0167`. For coding agents: read it all, follow the tasks in order, and check every acceptance criterion before you finish.

## Goal

Enable booking tickets for a group (up to 10) in one go instead of one person at a time, with a single confirmation to the booker and atomic all-or-nothing deduction from the 50-ticket pool.

## Decisions

- **One email confirmation to the booker per group, not per member**
- **Max group size is 10 people**
- **Group booking is atomic all-or-nothing against remaining tickets: reject whole group if total exceeds remaining**

## Open questions

Resolve these with the team before or while implementing; do not guess.

- Should groups get a discount? (explicitly undecided in conversation - do not implement)

## Code involved

| Path | Today | Change | Evidence |
|---|---|---|---|
| `main.go` | Infinite for{} loop books one person per iteration via getUserDetails, validates, mutates inventory, calls getTickets + go sendEmail. | Add group-size prompt/validation and member-name collection; enforce all-or-nothing deduction and single confirmation email. | `main.go:35-71`, `main.go:40-42`, `main.go:10-13`, `main.go:51-52`, `main.go:17-23`, `main.go:92-105` |
| `helper/helper.go` | IsValidateUserInputs checks name len>2, tickets>0 and <=remaining, email contains @, phone len==10. | Add group-size rule (1-10) alongside existing name/email/phone/ticket-count checks; reused for atomic group validation. | `main.go:40-42`, `helper/helper.go:7-12` |

## Approach

Follow existing main.go loop + helper pattern: keep fmt.Scan prompts in getUserDetails-style functions, keep UserData slice storage, keep single go sendEmail call per booking. Add group-size input (cap 10) validated before member collection, collect N member first/second names in one go, validate booker contact once with helper.IsValidateUserInputs, and check total group tickets <= remainingTicketsAvailable atomically before mutating ticketsSaled/remainingTicketsAvailable. No new frameworks or persistence.

## Tasks

1. **Add group-size validator** — In helper/helper.go add IsValidGroupSize(size uint) bool (1<=size<=10). Keep existing IsValidateUserInputs unchanged so single-booking rules still apply to booker.
   Files: `helper/helper.go`
2. **Add single-pass group input** — In main.go add group input flow: prompt booker for contact (first/second/email/phone) once + group size N + N member names. Validate booker with helper.IsValidateUserInputs and N with new validator. On failure print reason and continue without mutating inventory.
   Files: `main.go`
3. **Atomic inventory update + one confirmation** — Before any mutation check total group tickets (N, one ticket per member) <= remainingTicketsAvailable; if not, print shortage message and continue. On success do single ticketsSaled/remainingTicketsAvailable update, store group members, call getTickets once and go sendEmail once to booker email.
   Files: `main.go`

## Acceptance criteria

- [ ] Group booking of N<=10 with enough remaining tickets succeeds in one loop iteration and decrements remainingTicketsAvailable by total group tickets once
- [ ] Group with size >10 or size <=0 is rejected with message and remainingTicketsAvailable unchanged
- [ ] Group requesting more tickets than remaining is rejected whole (all-or-nothing) with remainingTicketsAvailable unchanged
- [ ] Successful group triggers exactly one sendEmail to booker email (no per-member emails)
- [ ] Invalid booker name/email/phone still rejected via existing IsValidateUserInputs rules

## Out of scope

- Group discounts/pricing changes
- Per-member emails or per-member ticket counts
