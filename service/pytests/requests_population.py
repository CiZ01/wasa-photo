#!/bin/python3

import requests as r
import json
import shutil
import random


username_list = {'mina', 'seulgi', 'joy', 'haru', 'jhope', 'sooyoung', 'ballo', 'jennie', 'yoona', 'yeri', 'rose', 'tzuyu', 'sana', 'seohyun', 'hyoyeon', 'jin', 'jessica', 'tiffany',
                 'nayeon', 'momo', 'jihyo', 'jungkook', 'jimin', 'yuri', 'wendy', 'sunny', 'irene', 'jeongyeon', 'dahyun', 'cha', 'suga', 'jisoo', 'lisa', 'chaeyoung', 'taeyeon', 'colla'}
randomCatPic = "https://cataas.com/cat"
randomQuote = "https://api.quotable.io/random?maxLength=64"


def create_users(count: int):
    for i in range(1,count):
        response = r.post("http://localhost:3000/session",
                          json={"username": username_list.pop()})
        if response.status_code != 201:
            print("Error creating user " + str(i))
            print(response.text)
            return


def create_posts(count: int):
    for i in range(1, count):
        for j in range(5):
            res = 0
            while res != 201:
                response_image = r.get(randomCatPic, stream=True)
                if response_image != '400':
                    with open("./tmp/tmp_img.jpeg", 'wb') as f:
                        response_image.raw.decode_content = True
                        shutil.copyfileobj(response_image.raw, f)

                response_quote = r.get(randomQuote)
                if response_quote != '400':
                    quote = response_quote.json()['content']
                header = {'Authorization': str(i)}
                multipart_form_data = {
                    'image': ('image.jpg', open('tmp/tmp_img.jpeg', 'rb')),
                    'caption': quote,
                }
                response = r.post(
                    f"http://localhost:3000/profiles/{i}/posts", headers=header, files=multipart_form_data)
                if response.status_code != 201:
                    print(f"Error creating post  user:{i} - post:{j}")
                    print(response.text)
                else:
                    res = response.status_code
                    print(f"{i}. Post succesfull posted!")


def create_likes(count: int) -> None:
    for i in range(1, count):
        header = {'Authorization': str((i%23)+1)}
        response = r.put(
            f"http://localhost:3000/profiles/21/posts/{random.randint(0,10)}/likes/{(i%23)+1}", headers=header)
        if response.status_code != 200:
            print("Error creating like " + str(i))
            print(response.text)
        else:
            print(f"{i}. Like succesfull posted!")


def create_comments(count):
    for i in range(1, count):
        text = r.get(randomQuote).json()['content']
        body = {
            'text': text
        }

        header = {'Authorization': str((i%23)+1)}
        response = r.post(
            f"http://localhost:3000/profiles/{random.randint(0,22)}/posts/{random.randint(0,5)}/comments", json=body, headers=header)
        if response.status_code != 201:
            print("Error creating comment " + str(i))
            print(response.text)
        else:
            print(f"{i}. Comment succesfull posted!")
    return


def create_follows(count: int):
    for i in range(0, count*2):
        user1, user2 = random.randint(1, 20), random.randint(1, 20)
        header = {'Authorization': str(user2)}
        response = r.put(
            f"http://localhost:3000/profiles/{user2}/followings/{user1}", headers=header)
        if response != '201':
            print(f"Follow non creato!, {user1} and 14")
            print(response.text)
        else:
            print(f"{i}. {response.text}")
    return


def create_ban(count: int):
    for i in range(1, count):
        user1, user2 = (random.randint(1, 20), random.randint(1, 20))
        header = {'Authorization': str(user1)}
        response = r.put(
            f"http://localhost:3000/profiles/{user2}/bans/{user1}", headers=header)
        if response != '200':
            print(f"Ban non creato!, {user1} and {user2}")
            print(response.text)
        else:
            print(f"{i}. {response.text}")
    return


def create_propic(count: int):
    for i in range(1, count):
        response_image = r.get(randomCatPic, stream=True)
        if response_image != '400':
            with open("./tmp/tmp_img.jpeg", 'wb') as f:
                response_image.raw.decode_content = True
                shutil.copyfileobj(response_image.raw, f)

        header = {'Authorization': str(i)}
        multipart_form_data = {
            'image': ('image.jpg', open('tmp/tmp_img.jpeg', 'rb')),
        }
        response = r.put(
            f"http://localhost:3000/profiles/{i}/profile-picture", headers=header, files=multipart_form_data)
        if response.status_code != 201:
            print(f"Error creating propic  user:{i}")
            print(response.text)
        else:
            print(f"{i}. Propic succesfull posted!")


def main():
    #create_users(20)
    #create_propic(20)
    #create_posts(22)
    create_likes(100)
    create_comments(30)
    create_follows(40)
    # create_ban(20)


if __name__ == '__main__':
    main()
