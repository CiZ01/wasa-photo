#!/bin/python3

import requests as r
import json
import shutil
import random


username_list = {'mina', 'seulgi', 'joy', 'haru', 'jhope', 'sooyoung', 'ballo', 'jennie', 'yoona', 'yeri', 'rose', 'tzuyu', 'sana', 'seohyun', 'hyoyeon', 'jin', 'jessica', 'tiffany',
                 'nayeon', 'momo', 'jihyo', 'jungkook', 'jimin', 'yuri', 'wendy', 'sunny', 'irene', 'jeongyeon', 'dahyun', 'cha', 'suga', 'jisoo', 'lisa', 'chaeyoung', 'taeyeon', 'colla'}
randomCatPic = "https://cataas.com/cat"
randomQuote = "https://api.quotable.io/random?maxLength=100"

def create_users(count: int):
    for i in range(count):
        response = r.post("http://localhost:3000/session",
                          json={"username":  username_list.pop()})
        if response.status_code != 201:
            print("Error creating user " + str(i))
            print(response.text)
            return


def create_posts(count: int):
    for i in range(1,count):
        for j in range(5):
            response_image = r.get(randomCatPic, stream=True)
            if response_image != '400':
                with open("./tmp/tmp_img.jpeg", 'wb') as f:
                    response_image.raw.decode_content = True
                    shutil.copyfileobj(response_image.raw, f)
            header = {'Authorization' : str(i)}
            multipart_form_data = {
                'image':('image.jpg', open('tmp/tmp_img.jpeg', 'rb')), 
                'caption':'sono haur!',
                }
            response = r.post(f"http://localhost:3000/profiles/{i}/posts", headers=header ,files=multipart_form_data)
            if response.status_code != 201:
                print(f"Error creating post  {i} - {j}")
                print(response.text)
            else:
                print(f"{i}. Post succesfull posted!")

def create_likes(count : int) -> None:
    i = 1

    while i < count:
        user1, user2 = (random.randint(1,20), random.randint(1,20))
        post_user2 = random.randint(2,5)
        header = {'Authorization': str(user1)}
        response = r.post(f"http://localhost:3000/profiles/{user2}/posts/{post_user2}/like", headers= header)
        if response != '201':
            print(f"Like.{i} non creato!, {user1} and {user2} and {post_user2}")
            print(response.text)
            i += 1
        else:
            print(f"{i}. {response.text}")
            i += 1

def create_comments(count):
    for i in range(1,count):
        text = r.get(randomQuote).json()['content']
        body = {
            'text': text
        }
        
        header = {'Authorization': str(i)}
        response = r.post(f"http://localhost:3000/profiles/{i}/posts/2/comments", json=body, headers=header)
        if response.status_code != 201:
            print("Error creating comment " + str(i))
            print(response.text)
        else:
            print(f"{i}. Comment succesfull posted!")
    return
        

def main():
    #create_users(20)
    #create_posts(20)
    create_likes(20)
    create_comments(30)


if __name__ == '__main__':
    main()
