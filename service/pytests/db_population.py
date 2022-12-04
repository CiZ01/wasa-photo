#!/user/bin/python3

import sqlite3
from sqlite3 import Error
import random
import time as time

db_file = '/tmp/wasa_TEST.db'

user = {
    'userID': 1,
    'username': 'haru',
    'userPropicURL': 'haru.jpg',
    'bio': 'NULL',
}

post = {
    'postID': 1,
    'userID': 1,
    'postImageURL': 'haru.jpg',
    'caption': 'NULL',
}

like = {
    'postID': 1,
    'userID': 1,
}

comment = {
    'commentID': 1,
    'postID': 1,
    'userID': 1,
    'commentText': 'Corona non perdona!',
}

ban = {
    'bannerID': 1,
    'bannedID': 2,
}

follow = {
    'followerID': 1,
    'followedID': 2,
}

username_list = ['mina', 'seulgi', 'joy', 'haru', 'jhope', 'sooyoung', 'ballo', 'jennie', 'yoona', 'yeri', 'rose', 'tzuyu', 'sana', 'seohyun', 'hyoyeon', 'jin', 'jessica', 'tiffany', 'nayeon', 'momo', 'jihyo', 'jungkook', 'jimin', 'yuri', 'rm', 'wendy', 'sunny', 'irene', 'v', 'jeongyeon', 'dahyun', 'cha', 'suga', 'jisoo', 'lisa', 'chaeyoung', 'taeyeon', 'colla']

def create_connection(db_file):
    conn = None
    try:
        conn = sqlite3.connect(db_file)
    except Error as e:
        print(e)

    return conn


def addUser(conn):
    sql = ' INSERT INTO User(userID,username,userPropicURL,bio) VALUES (?,?,?,?) '
    cur = conn.cursor()
    cur.execute(sql, (user['userID'], user['username'],
                user['userPropicURL'], user['bio']))
    conn.commit()
    return cur.lastrowid


def addPost(conn):
    sql = ' INSERT INTO Post(postID,userID,postImageURL,caption) VALUES (?,?,?,?) '
    cur = conn.cursor()
    cur.execute(sql, (post['postID'], post['userID'],
                post['postImageURL'], post['caption']))
    conn.commit()
    return cur.lastrowid

def addLike(conn):
    sql = ' INSERT INTO Like(postID,userID) VALUES (?,?) '
    cur = conn.cursor()
    cur.execute(sql, (like['postID'], like['userID']))
    conn.commit()
    return cur.lastrowid


def populateUser(conn):
    for i in range(1, len(username_list)):
        user['userID'] = i
        user['username'] = username_list[i]
        addUser(conn)


def populatePost(conn):
    for i in range(1, len(username_list)):
        post['postID'] = i
        if i % 2 == 0:
            post['userID'] = i
            addPost(conn)
        else:
            post['userID'] = i + 1
            addPost(conn)

def populateLike(conn):
    random.seed(time.time())
    for i in range(1, len(username_list)):
        like['postID'] = random.randint(1, len(username_list))
        like['userID'] = random.randint(1, len(username_list))
        addLike(conn)

def addComment(conn):
    sql = ' INSERT INTO Comment(postID,userID,commentText) VALUES (?,?,?) '
    cur = conn.cursor()
    cur.execute(sql, (comment['postID'], comment['userID'],
                comment['commentText']))
    conn.commit()
    return cur.lastrowid

def populateComment(conn):
    random.seed(time.time())
    for i in range(1, len(username_list)):
        comment['postID'] = random.randint(1, len(username_list))
        comment['userID'] = random.randint(1, len(username_list))
        comment['commentText'] = 'Corona non perdona!'
        addComment(conn)

def addBan(conn):
    sql = ' INSERT INTO Ban(bannerID,bannedID) VALUES (?,?) '
    cur = conn.cursor()
    cur.execute(sql, (ban['bannerID'], ban['bannedID']))
    conn.commit()
    return cur.lastrowid

def populateBan(conn):
    random.seed(time.time())
    for i in range(1, len(username_list)):
        ban['bannerID'] = random.randint(1, len(username_list))
        ban['bannedID'] = random.randint(1, len(username_list))
        addBan(conn)

def addFollow(conn):
    sql = ' INSERT INTO Follow(followerID,followedID) VALUES (?,?) '
    cur = conn.cursor()
    cur.execute(sql, (follow['followerID'], follow['followedID']))
    conn.commit()
    return cur.lastrowid

def populateFollow(conn):
    random.seed(time.time())
    for i in range(1, len(username_list)):
        follow['followerID'] = random.randint(1, len(username_list))
        follow['followedID'] = random.randint(1, len(username_list))
        if follow['followerID'] != follow['followedID']:
            addFollow(conn)

    
def main():
    tables_names = ['User', 'Post', 'Follow', 'Like', 'Comment', 'Ban']
    conn = create_connection(db_file)
    with conn:
        #populateUser(conn)
        #populatePost(conn)
        #populateLike(conn)
        #populateComment(conn)
        #populateBan(conn)
        populateFollow(conn)
        #clear_table(conn, tables_names[0])
        print_table(conn, tables_names[0])
    return 0

def clear_table(conn, table : str):
    sql = ' DELETE FROM ' + table + ';'
    cur = conn.cursor()
    cur.execute(sql)
    conn.commit()

def print_table(conn, table : str):
    sql = ' SELECT * FROM ' + table + ';'
    cur = conn.cursor()
    cur.execute(sql)
    rows = cur.fetchall()
    for row in rows:
        print(row)

if "__main__" == __name__:
    main()
    pass
