@board = [1, 2, 3, 4, 5, 6, 7, 8, 9]
@player1 = []
@player2 = []
@p1win = FALSE
@p2win = FALSE
@p1score = 0
@p2score = 0
@gameover = FALSE
@twoplayer = FALSE
@gamenumber = 0


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
  if @gameover == true
    puts "Game Completed! Player 1(X), what have you?"
    print "choose (#), or (n)ext game> "
    choice = gets.chomp
    if choice.upcase == "N" then
      go
    end
    move(choice.to_i)
  else
    puts "Player 1(X), what have you?"
    print "choose (#)> "
    choice = gets.chomp.to_i
    move(choice)
  end
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
  if @gameover == true
    puts "Game Completed! Player 2(O), what have you?"
    print "choose (#), or (n)ext game> "
    choice = gets.chomp
    if choice.upcase == "N" then
      go
    end
    move(choice.to_i)
  else
    puts "Player 2(O), what have you?"
    print "choose (#)> "
    choice = gets.chomp.to_i
    move(choice)
  end
end

def c_turn
  def move(n)
    @player2.push(n)
    @board.delete(n)
  end
  puts "Computer, what have you?"
  3.times do
    sleep 0.5
    print "."
  end
  sleep 1
  choice = @board.sample
  move(choice)
end

def p1_finnish(a, b, c)
  if @player1.include?(a) && @player1.include?(b) && @player1.include?(c)
    @p1win = true
    @gameover = true
    p1_wseq
  end
end

def p2_finnish(a, b, c)
  if @player2.include?(a) && @player2.include?(b) && @player2.include?(c)
    @p2win = true
    @gameover = true
    p2_wseq
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

def p1_wseq
  @p1score += 1
  3.times do
    @header = "           "
    board_moves
    sleep 0.5
    @header = "p1(X) wins!"
    board_moves
    sleep 0.5
  end
end

def p2_wseq
  @p2score += 1
  3.times do
    @header = "           "
    board_moves
    sleep 0.5
    if @twoplayer == true
      @header = "p2(O) wins!"
    elsif @twoplayer == false
      @header = "Comp wins!"
    end
    board_moves
    sleep 0.5
  end
end

def play_game
  until @board == nil
    p1_turn
    p1_win_check
    board_moves
    sleep 0.5

    if @twoplayer == true
      p2_turn
    else
      c_turn
    end
    p2_win_check
    board_moves
    sleep 0.5
  end
end

def reset_board
  @sq1 = "1"
  @sq2 = "2"
  @sq3 = "3"
  @sq4 = "4"
  @sq5 = "5"
  @sq6 = "6"
  @sq7 = "7"
  @sq8 = "8"
  @sq9 = "9"
  @header = "  BATTLE!  "
  @p1win = FALSE
  @p2win = FALSE
  @gameover = FALSE
  @board = [1, 2, 3, 4, 5, 6, 7, 8, 9]
  @player1 = []
  @player2 = []
  @gamenumber += 1
  3.times do
    sleep 0.5
    board_moves
    puts "Game ##{@gamenumber} go!"
    sleep 0.5
    board_moves
  end
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
  puts "    Tic-tac-toe    "
  puts "    #{@header}     "
  puts "╔════════╦════════╗"
  puts "║ X = #{@p1score}  ║ O = #{@p2score}  ║"
  puts "╚════════╩════════╝"
  puts "   ┍━━━━━━━━━━━┑"
  puts "   ⎟ #{@sq1} ║ #{@sq2} ║ #{@sq3} ⎟"
  puts "   ⎟═══╬═══╬═══⎟"
  puts "   ⎟ #{@sq4} ║ #{@sq5} ║ #{@sq6} ⎟"
  puts "   ⎟═══╬═══╬═══⎟"
  puts "   ⎟ #{@sq7} ║ #{@sq8} ║ #{@sq9} ⎟"
  puts "   ┕━━━━━━━━━━━┙"
end

def boot_game
  def to_start
    puts "Well, shall we do this?"
    print "(y)es or (n)o?> "
    start = gets.chomp
    if start.upcase == "Y"
      puts "LETS GET THIS PARTY STARTED"
      3.times do
        sleep 0.5
        # board_moves
        print "!"
        # board_moves
      end
    elsif start.upcase == "N"
      puts "Right"
      3.times do
        sleep 0.5
        print "."
      end
      print "so"
      sleep 3
      to_start
    else
      puts "YO! try that again!"
      sleep 1
      to_start
    end
  end
  def to_opponent
    puts "PvP or PvC"
    print "(c)omputer or (p)layer?> "
    opponent = gets.chomp
    if opponent.upcase == "C"
      @twoplayer = FALSE
      puts "ITS COMPUTER!"
      sleep 3
    elsif opponent.upcase == "P"
      @twoplayer = true
      puts "ITS ANOTHER PLAYER!"
      sleep 3
    else
      puts "YO! try that again!"
      sleep 1
      to_opponent
    end
  end
  to_start
  to_opponent
  @header = " BATTLE! "
end

def go
  reset_board
  play_game
end


boot_game
go
