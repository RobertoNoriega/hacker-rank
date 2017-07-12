#!/bin/python3

def compare_triplets(points_alice, points_bob):
    score_alice = 0
    score_bob = 0
    i = 0
    while i < len(points_alice):
        try:
            alice = int(points_alice[i])
            bob = int(points_bob[i])
            if alice > bob:
                score_alice = score_alice + 1 
            elif alice < bob:
                score_bob = score_bob + 1
            i = i + 1
        except Exception as ex:
            print(ex)
            break
    print(score_alice, score_bob)

points_alice = input().strip().split(" ")
points_bob = input().strip().split(" ")
if len(points_alice) == len(points_bob):
    compare_triplets(points_alice, points_bob)