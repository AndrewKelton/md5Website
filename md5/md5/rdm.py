from random import randint

def RCaptcha():
    rCap = []
    for i in range(7):
        rCap.append(randint(1,9)) 
    return rCap

if __name__ == "__main__":
    cap = RCaptcha()
    capS = ''.join(map(str, cap))
    print(capS)
