import turtle
import random

# Fungsi untuk menggambar kelopak bunga
def draw_petal():
    turtle.circle(100, 60)
    turtle.left(120)
    turtle.circle(100, 60)
    turtle.left(120)

# Fungsi untuk menggambar bunga lengkap dengan beberapa kelopak
def draw_flower():
    turtle.color("red")
    for _ in range(6):
        draw_petal()
        turtle.left(60)

# Fungsi untuk mengatur posisi turtle secara acak
def set_random_position():
    turtle.penup()
    x = random.randint(-200, 200)
    y = random.randint(-200, 200)
    turtle.goto(x, y)
    turtle.pendown()

# Menggambar beberapa bunga dengan posisi acak
for _ in range(5):
    set_random_position()
    draw_flower()

# Menutup jendela saat di-klik
turtle.exitonclick()
