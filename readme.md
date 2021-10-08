# spotrip

This is an open source Spotify song ripping application written in Golang
(using a modified version of
[Librespot](http://github.com/librespot-org/librespot-golang)).  It works with
both Spotify free and premium accounts.  It can download individual tracks,
albums, or entire artist discographies.   It downloads the audio files served
by Spotify's servers and does not work by recording audio outputs.

## Note

Use of this tool may result in your account getting banned by Spotify.
**Please** do not use it on accounts you have spent any money on.  I am not
liable for how you use this tool in any way.  From my experience using this
tool, I have observed that Spotify has rate limiting measures and must have
some idea of what is going on. 

## Dependencies

At least Go 1.11 to allow for module support.

## Build

There will be a warning about some undefined functions.  It can safely be
ignored.  This is because we have a local version of librespot, which adds
additional functionality.

## Usage

~~~
spotrip --username example@example.com --password password `
--tracks 4eJPzDcuWmXMedrwAbUeCt
~~~

## Additional Note on Search

To prevent yourself from generating a "suspicious" amount of logins from this
tool, you may want to use the search interface available at
[https://open.spotify.com/search](https://open.spotify.com/search). The ID
follows the last slash of the URL, for instance:

~~~
https://open.spotify.com/track/4qOPuARt2HNHAOlgXBezoT
~~~

## Credits

Thanks to [Librespot](https://github.com/librespot-org) for reverse engineering
the Spotify protocol and releasing their
[tools](https://github.com/librespot-org/spotify-analyze) for free. They are
the ones that make third party Spotify clients possible.
