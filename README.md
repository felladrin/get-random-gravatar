# Get Random Gravatar

PHP script to download a random gravatar for quick use.

For Linux users, just `chmod +x get_random_gravatar.php` and execute it.

For Windows users, just execute `get_random_gravatar.bat`.

Make sure you have `php` installed and set corretly in your System PATH.

As soon as you execute the script, a random gravatar will be downloaded and saved to the same folder.

You can change some configurations on the php file:

    $size_in_pixels = 96;
    $default_image = 'identicon';
    $rating = "g";

If you wanna know more about the configuration parameters, take a look at <http://gravatar.com/site/implement/images/>