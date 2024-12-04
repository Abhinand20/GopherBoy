import json
import os

OPCODE_DIR = "./Opcodes.json"
PREFIX = False

def format_instruction_record(key, value):
    """
    0xAB: func(cpu *CPU) {
        // MNEMONIC OP1, OP2
    },
    """
    actual_record = "{}: nil,\n".format(key)
    record = "{}: func(cpu *CPU) {{\n    ".format(key)
    mnemonic = value['mnemonic']
    operands = []
    for op in value['operands']:
        name = op['name']
        if op['immediate'] == False:
            name = "[{}]".format(name)
        if 'decrement' in op and op['decrement'] == True:
            mnemonic += "D"
        elif 'increment' in op and op['increment'] == True:
            mnemonic += "I"
        operands.append(name)
        
    comment = "// {} {}".format(
        mnemonic,
        ",".join(operands)
    )
    record += comment + "\n},"
    actual_record += "/*\n{}\n*/\n".format(record)
    instr_lookup_record = f'{key}: "{mnemonic} {",".join(operands)}",\n'
    return actual_record, instr_lookup_record
    
    
    
def main():
    with open(OPCODE_DIR, 'r') as f:
        data = json.load(f)
    
    prefixed = data['cbprefixed']
    if not PREFIX:
        prefixed = data['unprefixed']
    template = ""
    lookup_template = ""
    for k,v in prefixed.items():
        instr, lookup_instr = format_instruction_record(k,v)
        template += instr 
        lookup_template += lookup_instr
    with open("out.txt", 'w') as f:
        f.write(template)
    with open("out_lookup.txt", 'w') as f:
        f.write(lookup_template)
    
if __name__ == '__main__':
    main()