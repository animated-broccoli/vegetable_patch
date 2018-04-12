import praw
import pyimgur
import sys


if len(sys.argv) == 1:
    print("error")
    sys.exit()

if sys.argv[1] == "post":
    post()
elif sys.argv[1] == "fetch":
    
reddit = praw.Reddit(client_id='ZgqFkithftX0aw', client_secret="DL1KCB-fGgS583ezJNN0EXi6jS0",
                     password='kaleisgood', user_agent='MySimpleBot v0.1',
                     username='kalebot1')


client_id = "35819382c9c435c"
client_secret = "a39bd1bbae3dd782eed687a0befd3792582bd6d0"

im = pyimgur.Imgur(client_id)




def post():
    image = sys.argv[2]
    uploaded_image = im.upload_image(image, title="Check out this!")
    img_url = uploaded_image.link
    reddit.subreddit('animatedbroccoli').submit('Test Submission', url=img_url)



def fetch():
    submissions = reddit.subreddit('animatedbroccoli').new()
    for submission in submissions:
        print(submission.url)

fetch()

#post()


# import tweepy


# consumer_key = ""
# consumer_secret = ""
# access_token = ""
# access_token_secret = ""

# auth = tweepy.OAuthHandler(consumer_key, consumer_secret)
# auth.set_access_token(access_token, access_token_secret)

# api = tweepy.API(auth)

# public_tweets = api.home_timeline()
# for tweet in public_tweets:
#     print(tweet.text)