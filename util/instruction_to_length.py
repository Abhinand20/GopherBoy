
import json
import os

OPCODE_DIR = "./Opcodes.json"
    
    
def main():
    with open(OPCODE_DIR, 'r') as f:
        data = json.load(f)
    
    prefixed = data['unprefixed']
    template = "var instrLen = [0x100]byte{"
    i = 0
    for _,v in prefixed.items():
        if i > 0 and i % 16 == 0:
            template += "\n\t"
        template = template + str(v["bytes"]) + ", "
    with open("instr_len_map.txt", 'w') as f:
        f.write(template + "\n}")
    
if __name__ == '__main__':
    main()