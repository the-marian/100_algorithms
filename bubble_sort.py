from random import shuffle

l = list(range(1, 1000 + 1))
shuffle(l)
print(f"List before sorting {l} \n")
for i in range(1, len(l)):
    for j in range(len(l) - i):
        if l[j] > l[j + 1]:
            l[j], l[j + 1] = l[j + 1], l[j]
        else:
            continue

print(f"After bubble_sort {l}")