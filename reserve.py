import sys
theater, args = {row:[False]*5 for row in 'BDFHJ'}, sys.argv[1:]

def reserve(num_seats_requested:int, reserved: int) -> tuple:
    result = ''
    if num_seats_requested + reserved > 25: raise Exception()
    reserved += num_seats_requested
    for row in reversed(theater.keys()):
        if theater[row][4]: continue
        seat_num = 0
        while seat_num < len(theater[row]) and num_seats_requested > 0:
            if not theater[row][seat_num]:
                theater[row][seat_num] = True
                result += f'{row}{4*seat_num+1} '
                num_seats_requested-=1
            seat_num += 1
    return result, reserved

def main():
    if len(args) == 0: raise Exception('input file missing!')
    input_file_name, output_file_name, reserved = args[0], args[1] if len(args) == 2 else 'output.txt', 0
    in_file, out_file = open(input_file_name, 'r'), open(output_file_name, 'w')
    for line in in_file.readlines():
        request = line.split()
        try:
            result, reserved = reserve(int(request[1]), reserved)
        except:
            print(f'Request {request[0]} could not be supported: Theater has no space to support this request.')
        else:
            print(f'{request[0]} {result}', file=out_file)
    in_file.close()
    out_file.close()

if __name__ == "__main__":
    main()  

