# Spotify

This is an open source Spotify song ripping application written in Golang. It
works with both Spotify free and premium accounts.  It can download individual
tracks, albums, or entire artist discographies.   It downloads the audio files
served by Spotify's servers and does not work by recording audio outputs.

## Note

Use of this tool may result in your account getting banned by Spotify.
**Please** do not use it on accounts you have spent any money on.  I am not
liable for how you use this tool in any way.  From my experience using this
tool, I have observed that Spotify has rate limiting measures and must have
some idea of what is going on. 

## Dependencies

At least Go 1.11 to allow for module support.

## Usage

~~~
spotify --username example@example.com --password password `
--tracks 4eJPzDcuWmXMedrwAbUeCt
~~~

## Additional Note on Search

To prevent yourself from generating a "suspicious" amount of logins from this
tool, you may want to use the search interface available at
[https://open.spotify.com/search](https://open.spotify.com/search).  Simply
right click the album, track, or artist you wish to download and press "Copy
Song Link." 

The ID follows the last slash of the URL, for instance, in:

~~~
https://open.spotify.com/track/4qOPuARt2HNHAOlgXBezoT
~~~

`4qOPuARt2HNHAOlgXBezoT` would be the track ID.
