from collections import defaultdict

filename = 'input.txt'

# We will keep track of everything in a dictionary
# Since we don't need the full polymer, we just note the
# pairs, elements, and the count of each

rules = {}                          # E.g rule AB -> C: key=AB, value=C
element_pairs = defaultdict(int)    # E.g key=AB, value=integer
count_dict = defaultdict(int)       # E.g key=A, value=integer


# Process input
with open(filename) as file:
    polymer_template = file.readline().strip()

    for line in file:
        if line.strip() != '':    # Skip the blank line
            pair, element = line.strip().split(' -> ')
            rules[pair] = element

for c in range(0, len(polymer_template) - 1):
    this_pair = polymer_template[c] + polymer_template[c + 1]
    element_pairs[this_pair] += 1
    count_dict[polymer_template[c]] += 1

count_dict[polymer_template[-1]] += 1


def do_step(current_pairs, element_count, insertion_rules):
    """A function to insert elements between pairs according to the insertion
    rules. For each pair, it adds 2 new resulting pairs to a new dictionary
    of pairs. E.g. CD > E results in new pairs CE and ED. If a pair in the
    current_pair dict doesn't have an insertion rule, it just copies that pair
    to the new dict. This function also updates the count. E.g., for CD > E,
    update the count of E by the count of pair CD."""

    new_pairs = defaultdict(int)
    for the_pair in current_pairs:
        if the_pair in insertion_rules:
            new_element = insertion_rules[the_pair]
            new1 = the_pair[0] + new_element
            new2 = new_element + the_pair[1]
            new_pairs[new1] += current_pairs[the_pair]
            new_pairs[new2] += current_pairs[the_pair]
            element_count[new_element] += current_pairs[the_pair]
        else:
            new_pairs[the_pair] = current_pairs[the_pair]
    return new_pairs


steps = 40

for i in range(0, steps):
    element_pairs = do_step(element_pairs, count_dict, rules)

print(max(count_dict.values()) - min(count_dict.values()))
