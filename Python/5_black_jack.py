# Instructions
# In this exercise you are going to implement some rules of Blackjack, such as the way the game is played and scored.

# Note : In this exercise, A means ace, J means jack, Q means queen, and K means king. Jokers are discarded. A standard French-suited 52-card deck is assumed, but in most versions, several decks are shuffled together for play.

# 1. Calculate the value of a card
# In Blackjack, it is up to each individual player if an ace is worth 1 or 11 points (more on that later). Face cards (J, Q, K) are scored at 10 points and any other card is worth its "pip" (numerical) value.

# Define the value_of_card(<card>) function with parameter card. The function should return the numerical value of the passed-in card string. Since an ace can take on multiple values (1 or 11), this function should fix the value of an ace card at 1 for the time being. Later on, you will implement a function to determine the value of an ace card, give an existing hand.

# >>> value_of_card('K')
# 10

# >>> value_of_card('4')
# 4

# >>> value_of_card('A')
# 1
# 2. Determine which card has a higher value
# Define the higher_card(<card_one>, <card_two>) function having parameters card_one and card_two. For scoring purposes, the value of J, Q or K is 10. The function should return which card has the higher value for scoring. If both cards have an equal value, return both. Returning both cards can be done by using a comma in the return statement:

# # Using a comma in a return creates a Tuple.  Tuples will be covered in a later exercise.
# >>> def returning_two_values(value_one, value_two):
#         return value_one, value_two

# >>> returning_two_values('K', '3')
# ('K', '3')
# An ace can take on multiple values, so we will fix A cards to a value of 1 for this task.

# >>> higher_card('K', '10')
# ('K', '10')

# >>> higher_card('4', '6')
# '6'

# >>> higher_card('K', 'A')
# 'K'
# 3. Calculate the value of an ace
# As mentioned before, an ace can be worth either 1 or 11 points. Players try to get as close as possible to a score of 21, without going over 21 (going "bust").

# Define the value_of_ace(<card_one>, <card_two>) function with parameters card_one and card_two, which are a pair of cards already in the hand before getting an ace card. Your function will have to decide if the upcoming ace will get a value of 1 or a value of 11, and return that value. Remember: the value of the hand with the ace needs to be as high as possible without going over 21.

# Hint: if we already have an ace in hand then it's value would be 11.

# >>> value_of_ace('6', `K`)
# 1

# >>> value_of_ace('7', '3')
# 11
# 4. Determine a "Natural" or "Blackjack" Hand
# If the first two cards a player is dealt are an ace (A) and a ten-card (10, K, Q or J), giving a score of 21 in two cards, the hand is considered a natural or blackjack.

# Define the is_blackjack(<card_one>, <card_two>) function with parameters card_one and card_two, which are a pair of cards. Determine if the two-card hand is a blackjack, and return the boolean True if it is, False otherwise.

# Note : The score calculation can be done in many ways. But if possible, we'd like you to check if there is an ace and a ten-card in the hand (or at a certain position), as opposed to summing the hand values.

# >>> is_blackjack('A', 'K')
# True

# >>> is_blackjack('10', '9')
# False
# 5. Splitting pairs
# If the players first two cards are of the same value, such as two sixes, or a Q and K a player may choose to treat them as two separate hands. This is known as "splitting pairs".

# Define the can_split_pairs(<card_one>, <card_two>) function with parameters card_one and card_two, which are a pair of cards. Determine if this two-card hand can be split into two pairs. If the hand can be split, return the boolean True otherwise, return False

# >>> can_split_pairs('Q', 'K')
# True

# >>> can_split_pairs('10', 'A')
# False
# 6. Doubling down
# When the original two cards dealt total 9, 10, or 11 points, a player can place an additional bet equal to their original bet. This is known as "doubling down".

# Define the can_double_down(<card_one>, <card_two>) function with parameters card_one and card_two, which are a pair of cards. Determine if the two-card hand can be "doubled down", and return the boolean True if it can, False otherwise.

# >>> can_double_down('A', '9')
# True

# >>> can_double_down('10', '2')
# False """

# """Functions to help play and score a game of blackjack.

# How to play blackjack:    https://bicyclecards.com/how-to-play/blackjack/
# "Standard" playing cards: https://en.wikipedia.org/wiki/Standard_52-card_deck



def value_of_card(card):
    """Determine the scoring value of a card.

    :param card: str - given card.
    :return: int - value of a given card. 'J', 'Q', 'K' = 10; 'A' = 1; numerical value otherwise.
    """

    if card == 'J' or card == 'K' or card == 'Q':
        return 10
    elif card == 'A':
        return 1
    else:
        return "Error: invalid input"


def higher_card(card_one, card_two):
    """Determine which card has a higher value in the hand.

    :param card_one, card_two: str - cards dealt. 'J', 'Q', 'K' = 10; 'A' = 1; numerical value otherwise.
    :return: higher value card - str. Tuple of both cards if they are of equal value.
    """

    if card_one > card_two:
        return card_one
    elif card_one == card_two:
        return card_one, card_two
    else:
        return card_two


def value_of_ace(card_one, card_two):
    """Calculate the most advantageous value for the ace card.

    :param card_one, card_two: str - card dealt. 'J', 'Q', 'K' = 10; 'A' = 11 (if already in hand); numerical value otherwise.
    :return: int - value of the upcoming ace card (either 1 or 11).
    """

    value_one = value_of_card(card_one)
    value_two = value_of_card(card_two)

    if value_one + value_two + 11 < 21:
        return 1
    else:
        return 11


def is_blackjack(card_one, card_two):
    """Determine if the hand is a 'natural' or 'blackjack'.

    :param card_one, card_two: str - cards dealt. 'J', 'Q', 'K' = 10; 'A' = 11; numerical value otherwise.
    :return: bool - if the hand is a blackjack (two cards worth 21).
    """

    pass


def can_split_pairs(card_one, card_two):
    """Determine if a player can split their hand into two hands.

    :param card_one, card_two: str - cards dealt.
    :return: bool - if the hand can be split into two pairs (i.e. cards are of the same value).
    """

    pass


def can_double_down(card_one, card_two):
    """Determine if a blackjack player can place a double down bet.

    :param card_one, card_two: str - first and second cards in hand.
    :return: bool - if the hand can be doubled down (i.e. totals 9, 10 or 11 points).
    """

    pass



#1.
print(value_of_card('J'))
print(value_of_card('K'))
print(value_of_card('A'))

#2.
print(higher_card('K', '10'))
print(higher_card('4', '6'))
print(higher_card('K', 'A'))

#3.
print(value_of_ace('6', 'K'))
print(value_of_ace('7', '3'))