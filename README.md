back-finder
============

Lib to navigate up the folder tree to find a file with a given file name.

Why ?
-----
For scripting purpose mostly. That's a little like `git` works, actually. It detects that there is a `.git`
folder somewhere in a upper directory and allow you to use `git` commands. If the folder does not exist, that simply means
that you are not inside a `git` tracked folder, and you won't be able to perform most of `git` actions.

I often have this use case when I need to find a specific file (or check if it exists) in a upper directory for script purpose.
Additionally, I wanted an easy-to-use solution instead of relying on the same bash script block that only make script
harder to read and manage.

Originally it was a stand-alone application. But since this is a feature that I need to integrate in other application,
it has been extracted into a dedicated lib.
