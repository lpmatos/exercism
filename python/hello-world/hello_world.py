from typing import Text

def hello(name: Text = "World") -> Text:
    return u"Hello, {name}!".format(name=name)
