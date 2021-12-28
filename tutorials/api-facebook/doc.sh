# API v12.0
# API call
https://graph.facebook.com/USER-ID?access_token=ACCESS-TOKEN

# Get user ID,Name
https://graph.facebook.com/me?fields=id,name&access_token=TOKEN

# Get all feeds
https://graph.facebook.com/v12.0/me/feed?fields=id,name,message,story,created_time,link,description,caption,attachments{media,type,subattachments}&limit=100&access_token=TOKEN&format=json

# Get all posts
"https://graph.facebook.com/v12.0/me?fields=id,name,posts&access_token=TOKEN"
