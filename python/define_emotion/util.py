import random

alph = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"


def gen_rand_description(descLen: int) -> str:
    description = ""
    for i in range(descLen):
        description += alph[random.randint(0, len(alph)-1)]
    return description
