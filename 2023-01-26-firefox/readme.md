# Firefox

- https://archive.mozilla.org/pub/firefox/releases/
- https://archive.mozilla.org/pub/mobile/releases/68.11.0/

## bookmarklet

~~~js
javascript: {
   const n = 10;
   console.log(n);
} void 0;
~~~

## policies.json

~~~
C:\Program Files\Mozilla Firefox\distribution\policies.json
~~~

## resetUserPrefs

<https://bugzilla.mozilla.org/show_bug.cgi?id=1687196>

## command options

Bypass Profile Manager and launch application with a profile:

~~~
waterfox -no-remote -P May
~~~

Start with Profile Manager:

~~~
waterfox -P
~~~

Make sure to not name the default profile, as then script will require:

~~~
waterfox -P May
~~~

If you did this already, just rename to a blank name.

<https://developer.mozilla.org/Mozilla/Command_Line_Options>
