
@p1score = 0
@p2score = 0
@gameover = FALSE
@twoplayer = FALSE
@gamenumber = 0

def move(player)
  choice = gets.chomp.to_i
  if @board.include?(choice)
    player.push(choice)
    @board.delete(choice)
  else
    puts "Play not available, please try again!"
    print "> "
    move(player)
  end
end

def continue
  puts "Shall we play again?"
  puts "(n)ext game, (r)eturn to menu, or (q)uit"
  print "> "
  choice = gets.chomp
  case choice.upcase
  when "Q"
    exit
  when "N"
    @gameover = false
  when "R"
    @gameover = true
    @p1score = 0
    @p2score = 0
  else
    puts "I'm sorry, we didn't get that - lets try again."
    sleep 1
    board_moves
    continue
  end
end

def p1_turn
  player = @player1
  puts "Player 1(X), what have you?"
  print "> "
  move(player)
end

def p2_turn
  player = @player2
  puts "Player 2(O), what have you?"
  print "> "
  move(player)
end

def c_turn
  def computer_move(choice)
    @player2.push(choice)
    @board.delete(choice)
  end
  puts "Computer, what have you?"
  timer("period", 3, 1)
  choice = @board.sample
  computer_move(choice)
end

def p1_finish(a, b, c)
  if @player1.include?(a) && @player1.include?(b) && @player1.include?(c)
    @p1win = true
    @gameover = true
    p1_win_sequence
  end
end

def p2_finish(a, b, c)
  if @player2.include?(a) && @player2.include?(b) && @player2.include?(c)
    @p2win = true
    @gameover = true
    p2_win_sequence
  end
end

def p1_win_check
  unless @p1win || @p2win
    p1_finish(1,2,3)
    p1_finish(4,5,6)
    p1_finish(7,8,9)
    p1_finish(1,4,7)
    p1_finish(2,5,8)
    p1_finish(3,6,9)
    p1_finish(1,5,9)
    p1_finish(3,5,7)
  end
end

def p2_win_check
  unless @p1win || @p2win
    p2_finish(1,2,3)
    p2_finish(4,5,6)
    p2_finish(7,8,9)
    p2_finish(1,4,7)
    p2_finish(2,5,8)
    p2_finish(3,6,9)
    p2_finish(1,5,9)
    p2_finish(3,5,7)
  end
end

# def flash_header(blink_words, number_of)
#   number_of.times do
#     @header = "           "
#     board_moves
#     sleep 0.5
#     @header = "#{blink_words}"
#     board_moves
#     sleep 0.5
#   end
# end

def p1_win_sequence
  @p1score += 1
  3.times do
    @header = "           "
    board_moves
    sleep 0.5
    @header = "p1(X) wins!"
    board_moves
    sleep 0.5
  end
  go
end

def p2_win_sequence
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
  go
end

def cat_win_sequence
  3.times do
    @header = "           "
    board_moves
    sleep 0.5
    @header = " TIE-GAME! "
    board_moves
    sleep 0.5
  end
  go
end

def play_game
  unless @board == []
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
  else
    cat_win_sequence
  end
  play_game
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

# i = 0
#
# 9.times do(i)
#   i += 1
#   puts i
# end
#
#   end

def board_moves
  x_marks
  o_marks
  system("clear")
  unless 1 == 0
    puts "    Tic-tac-toe    "
    puts "    #{@header}     "
    puts "╔════════╦════════╗"
    puts "║ X = #{@p1score}  ║ O = #{@p2score}  ║"
    puts "╚════════╩════════╝"
  end
  unless 1 == 0
    puts "   ┍━━━━━━━━━━━┑"
    puts "   ⎟ #{@sq1} ║ #{@sq2} ║ #{@sq3} ⎟"
    puts "   ⎟═══╬═══╬═══⎟"
    puts "   ⎟ #{@sq4} ║ #{@sq5} ║ #{@sq6} ⎟"
    puts "   ⎟═══╬═══╬═══⎟"
    puts "   ⎟ #{@sq7} ║ #{@sq8} ║ #{@sq9} ⎟"
    puts "   ┕━━━━━━━━━━━┙"
  end
  puts ""
end

def boot_game
  def intro(try_again)
    system('clear')
    print "TIC-TAC-TOE!"
    timer("period", 3, 1) unless try_again == true
    print "LETS GET THIS PARTY STARTED! "
    timer("blank", 2, 0) unless try_again == true
    puts "PvP or PvC?"
  end

  def pick
    puts "pick (c)omputer or (p)layer?... or just (q)"
    print "> "
    answer = gets.chomp
    case answer.upcase
    when "C"
      @twoplayer = FALSE
      system('clear')
      print "ITS COMPUTER!"
      timer("exclaim", 2, 1)
    when "P"
      @twoplayer = true
      system('clear')
      print "ITS ANOTHER PLAYER!"
      timer("exclaim", 2, 1)
    when "Q"
      system('clear')
      print "Well, enjoy your day then!"
      timer("blank", 2, 1)
      exit
    else
      puts "Yo, you're answer is not valid homeboy - try that again!"
      timer("blank", 2, 1)
      intro(try_again = true)
      pick
    end
  end

  intro(try_again = false)
  pick
  @header = " BATTLE! "
end

def timer(action, number_of, extra_time)
  number_of.times do
    sleep 0.5
    if action == "exclaim" then print "!" end
    if action == "period" then print "." end
    if action == "blank" then print "" end
  end
  sleep extra_time
end

def go
  reset_board
  play_game
end

boot_game
go
