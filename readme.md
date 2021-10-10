# Spotify

This module works with both Spotify free and premium accounts. It can download
individual tracks, albums, or entire artist discographies. It downloads the
audio files served by Spotify's servers and does not work by recording audio
outputs.

## Note

Use of this tool may result in your account getting banned by Spotify.
**Please** do not use it on accounts you have spent any money on. I am not
liable for how you use this tool in any way. From my experience using this
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
tool, you may want to use the search interface available at Spotify. Simply
right click the album, track, or artist you wish to download and click Share,
Copy Song Link.
