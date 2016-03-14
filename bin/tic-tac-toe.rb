@board = [1, 2, 3, 4, 5, 6, 7, 8, 9]
@player1 = []
@player2 = []
@p1win = FALSE
@p2win = FALSE


def p1_turn
  def move(n)
    if @board.include?(n)
      @player1.push(n)
      @board.delete(n)
    else
      puts "Play not available, please try again!"
      p1_turn
    end
  end
  puts "Player 'X', what have you?"
  choice = gets.chomp.to_i
  move(choice)
end

def p2_turn
  def move(n)
    if @board.include?(n)
      @player2.push(n)
      @board.delete(n)
    else
      puts "Play not available, please try again!"
      p2_turn
    end
  end
  puts "Player 'O', what have you?"
  choice = gets.chomp.to_i
  move(choice)
end

def p1_finnish(a, b, c)
  if @player1.include?(a) && @player1.include?(b) && @player1.include?(c)
    @p1win = true
    @header = "X wins!"
    p2_turn
  end
end

def p2_finnish(a, b, c)
  if @player2.include?(a) && @player2.include?(b) && @player2.include?(c)
    @p2win = true
    @header = "O wins!"
    p1_turn
  end
end

def p1_win_check
  unless @p1win || @p2win
    p1_finnish(1,2,3)
    p1_finnish(4,5,6)
    p1_finnish(7,8,9)
    p1_finnish(1,4,7)
    p1_finnish(2,5,8)
    p1_finnish(3,6,9)
    p1_finnish(1,5,9)
    p1_finnish(3,5,7)
  end
end

def p2_win_check
  unless @p1win || @p2win
    p2_finnish(1,2,3)
    p2_finnish(4,5,6)
    p2_finnish(7,8,9)
    p2_finnish(1,4,7)
    p2_finnish(2,5,8)
    p2_finnish(3,6,9)
    p2_finnish(1,5,9)
    p2_finnish(3,5,7)
  end
end

def play_game
  until @board == []
    p1_turn
    p1_win_check
    board_moves

    p2_turn
    p2_win_check
    board_moves

  end
end



def set_board
  @sq1 = "1"
  @sq2 = "2"
  @sq3 = "3"
  @sq4 = "4"
  @sq5 = "5"
  @sq6 = "6"
  @sq7 = "7"
  @sq8 = "8"
  @sq9 = "9"
  @header = " ready "
end

def x_marks
  if @player1.include?(1) then
    @sq1 = "X"
  end
  if @player1.include?(2) then
    @sq2 = "X"
  end
  if @player1.include?(3) then
    @sq3 = "X"
  end
  if @player1.include?(4) then
    @sq4 = "X"
  end
  if @player1.include?(5) then
    @sq5 = "X"
  end
  if @player1.include?(6) then
    @sq6 = "X"
  end
  if @player1.include?(7) then
    @sq7 = "X"
  end
  if @player1.include?(8) then
    @sq8 = "X"
  end
  if @player1.include?(9) then
    @sq9 = "X"
  end
end

def o_marks
  if @player2.include?(1) then
    @sq1 = "O"
  end
  if @player2.include?(2) then
    @sq2 = "O"
  end
  if @player2.include?(3) then
    @sq3 = "O"
  end
  if @player2.include?(4) then
    @sq4 = "O"
  end
  if @player2.include?(5) then
    @sq5 = "O"
  end
  if @player2.include?(6) then
    @sq6 = "O"
  end
  if @player2.include?(7) then
    @sq7 = "O"
  end
  if @player2.include?(8) then
    @sq8 = "O"
  end
  if @player2.include?(9) then
    @sq9 = "O"
  end
end

def board_moves
  system("clear")
  x_marks
  o_marks
  puts " Tic-tac-toe "
  puts "   #{@header}   "
  puts "┍━━━━━━━━━━━┑"
  puts "⎟ #{@sq1} ║ #{@sq2} ║ #{@sq3} ⎟"
  puts "⎟═══╬═══╬═══⎟"
  puts "⎟ #{@sq4} ║ #{@sq5} ║ #{@sq6} ⎟"
  puts "⎟═══╬═══╬═══⎟"
  puts "⎟ #{@sq7} ║ #{@sq8} ║ #{@sq9} ⎟"
  puts "┕━━━━━━━━━━━┙"
end

def boot_game
  board_moves
  puts "Well, shall we do this?"
  print "(y)es or (n)o?> "
  start = gets.chomp
  if start.upcase == "Y"
    3.times do
      sleep 0.5
      board_moves
      puts "LETS GET THIS PARTY STARTED!"
      sleep 0.5
      board_moves
    end
  elsif start.upcase == "N"
    board_moves
    puts "Right"
    3.times do
      sleep 1
      print "."
    end
    print "so"
    sleep 3
    boot_game
  else
    board_moves
    puts "YO! try that again!"
    sleep 1
    boot_game
  end
  @header = "BATTLE!"
end

set_board
boot_game
play_game
board_moves
