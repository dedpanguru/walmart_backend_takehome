from math import remainder
import sys
theater, args = {row:[False]*5 for row in 'BDFHJ'}, sys.argv[1:]

def reserve_group_in_same_row(num_seats_requested: int, reserved: int) -> tuple:
    """
        tries to seat the request in the same row
        params:
            num_seats_requested -> integer representing the number of requested seats
            reserved -> integer representing the number of seats already reserved
        return:
            tuple holding string representing reservations made, the amount of seats now reserved, and a boolean variable indicating success of the reservations
    """
    reservation, made_reservation = '', False
    if num_seats_requested + reserved > 25: raise Exception
    reserved += num_seats_requested
    for row in reversed(theater.keys()):
        if theater[row].count(False) >= num_seats_requested:
            seat_num = theater[row].index(False)
            while seat_num <= len(theater[row])-1 and num_seats_requested >0:
                theater[row][seat_num] = True
                reservation += f'{row}{4*seat_num+2} '
                seat_num+=1
                num_seats_requested-=1
            made_reservation = True
    return reservation, reserved, made_reservation

def reserve_individual(num_seats_requested:int, reserved: int) -> tuple:
    """
    Allocates the number of seats requested in num_seats_requested in the theater to the request
    params:
        num_seats_requested -> integer representing the number of seats requested
        reserved -> integer representing number of seats already reserved
    return:
        tuple holding the string representing the reservation and the number of seats reserved
    """
    result = '' # will hold the reservation string
    # insufficient space check
    if num_seats_requested + reserved > 25: raise Exception
    reserved += num_seats_requested # sufficient space means that all request seats will be reserved, so updated reserved varaiable
    for row in reversed(theater.keys()): # starting from the back rows
        if not theater[row].count(False): continue # skip the row if it has already been reserved
        seat_num = 0
        while seat_num < len(theater[row]) and num_seats_requested > 0:
            if not theater[row][seat_num]:
                theater[row][seat_num] = True
                result += f'{row}{4*seat_num+2} '
                num_seats_requested-=1
            seat_num += 1
    #print(reserved)
    return result, reserved

def main():
    # command-line argument check
    if len(args) == 0: raise Exception('input file missing!')
    # access command-line arguments, default output file name is 'output.txt'
    input_file_name, output_file_name = args[0], args[1] if len(args) == 2 else 'output.txt'
    reserved = 0
    #open both files
    in_file, out_file = open(input_file_name, 'r'), open(output_file_name, 'w')
    # scan each line in the input file
    for line in in_file.readlines():
        request = line.split()
        num_seats_requested = int(request[1])
        # try to reserve seats
        try:
            # prioritize grouping in the same row
            reservation, reserved, worked = reserve_group_in_same_row(num_seats_requested,reserved)
            if not worked: # the request cannot be grouped in the same row, reservation = '' on failure
                reserved -= num_seats_requested # delete the recorded reservations from reserved
                if num_seats_requested > 5: # if the row cannot hold the request
                    # split the request up into multiple requests of 5, and seat the offset individually
                    rounds, remainder= num_seats_requested//5, num_seats_requested%5 
                    for _ in range(rounds): # fill available rows 
                        temp, reserved, worked = reserve_group_in_same_row(5, reserved)
                        if worked:
                           reservation += temp
                        else:
                            reserved -= 5
                    num_seats_requested = remainder # after filling whole available rows, focus on seats that are left 
                temp, reserved = reserve_individual(num_seats_requested, reserved)
                reservation += temp
        except:
            # print out request reject message on error
            print(f'Request {request[0]} rejected: Theater has no space to support a request of {request[1]} seats. Theater capacity = {sum(seats.count(False)for seats in theater.values())}')
        else:
            # write to output file if there is no error
            print(f'{request[0]} {reservation}', file=out_file)
    # close both files once finished
    in_file.close()
    out_file.close()

if __name__ == "__main__":
    main()  
