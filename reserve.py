import theater
import sys

args = sys.argv[1:]
if len(args) == 0: raise Exception('input file missing!')
input_file_name = args[0]
output_file_name = args[1] if len(args) == 2 else 'output.txt'
reserved = 0
with open(input_file_name,'r') as in_file:
    with open(output_file_name, 'w') as out_file:
        for line in in_file.readlines():
            request = line.split()
            try:
                result, reserved = theater.reserve(int(request[1]), reserved if reserved else 0)
            except:
                print(f'Request {request[0]} could not be supported: Theater has no space to support this request.')
            else:
                print(f'{request[0]} {result}', file=out_file)
                

