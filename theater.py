theater = {row:[False]*5 for row in 'BDFHJ'}
def reserve(num_seats_requested:int, reserved: int)->tuple:
    result = ''
    if num_seats_requested + reserved > 25: raise Exception('insufficient space in theater!')
    reserved += num_seats_requested
    for row in reversed(theater.keys()):
        if theater[row][4]: continue
        seat_num = 0
        while seat_num < len(theater[row]) and num_seats_requested > 0:
            if not theater[row][seat_num]:
                theater[row][seat_num] = True
                result+=f'{row}{4*seat_num+1} '
                num_seats_requested-=1
            seat_num+=1
    #print(theater)
    return result,reserved