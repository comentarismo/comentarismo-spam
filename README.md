Comentarismo Spam 

# A API to stop spam on comments

# Options
```
SPAM_DEBUG, if true will debug all log entries for spam detection (optional)

LEARNSPAM, if true will use config_spamwords for each available lang, for testing proposes only (optional)
REDIS_HOST, ip addr of the redis instance to be used (required) -> defaults to g7-host
REDIS_PORT, port number of the redis instance to be used (required) defaults to 6379
REDIS_PASSWORD, password for this instance to be used (optional)
```

Running with defaults, to start in debug mode and learn spammy words
```
$ SPAM_DEBUG=true LEARNSPAM=true godep go run main.go
```

```
$ SPAM_DEBUG=true LEARNSPAM=true godep go run main.go
```



Testing API Spamassasssin
```
$ SPAM_DEBUG=false LEARNSPAM=true go test server_sa_test.go -v
```



```
sudo aptitude install spamassassin spamc
```
english spam words
http://help.pardot.com/customer/portal/articles/2126040-spam-words-reference
https://emailmarketing.comm100.com/email-marketing-ebook/spam-words.aspx
http://www.mailup.com/resources-mailup/strategy/strategies-techniques-and-best-practices/spam-words-to-avoid/?pd=80793116c3ae1
http://blog.hubspot.com/blog/tabid/6307/bid/30684/The-Ultimate-List-of-Email-SPAM-Trigger-Words.aspx#sm.00000ufaikg17y4e0uum3qchh7z5j


http://www.bannedwordlist.com/lists/swearWords.txt


https://www.npmjs.com/package/honeypot
http://stopforumspam.com/

A DNSBL server and RBL Checker client written in node.js
https://github.com/jawsome/node-dnsbl

https://www.npmjs.com/browse/keyword/spam?offset=36

http://www.trevorayers.com/wordpress-spam-filter-keyword-list/


```
Word,Score
daily!,571
pledge,341
forget,225
winner's,219
warranty,163
winners,160
analyses,121
affiliate,119
url,89
anti,86
annoying,82
for!,75
respect,73
private,71
huh,71
insurance,69
prior,68
biz,67
headaches,65
freebies,64
redirected,63
fashion,62
joining,59
gevalia,57
state,54
robe,53
cards,53
oscar,53
programmable,52
scratch,52
mascara,52
suite,52
vote!,51
eye,51
location,51
messages,51
specific,50
general,50
send,49
directories,48
end,48
emails,48
market,47
delay,47
programs,46
friends!,46
with,45
useful,45
career,45
presented,45
reserved,44
evans,44
minimum,44
receive,44
count,43
job,43
else,43
totals,43
looking,43
webmaster,43
actual,43
interested,43
cliquez,42
bank,42
credit,42
serve,41
towards,41
marketer,41
join,41
either,41
for,41
require,41
red,40
florida,40
tools,40
earning,40
opportunities,40
discontinue,39
advantage,39
simply,39
links,39
anytime,39
west,39
fit,39
earned,39
direct,38
tour,38
provided,38
check,38
attention,38
benefit,38
let,37
deck,37
moisture,36
limited,36
mayer,36
list,36
rights,35
corey,35
update,35
charge,35
card,35
afford,34
north,34
ion,34
currently,34
well,34
vary,34
clicking,34
we'll,34
does,34
steps,33
services,33
millions,33
hello,33
way,33
resources,33
included,33
they're,33
owners,33
take,32
expenses,32
joined,32
unique,32
concerns,32
you've,32
special,32
exciting,32
barton,32
entitles,32
months,32
next,32
updated,32
frustration,32
e-book,32
custom,31
support,31
prepare,31
costs,31
mugs,31
everyone!,31
don't,31
code,31
bracelet,31
feel,31
who,31
information,30
hotel,30
new,30
position,30
kids,30
practically,30
watch,30
effort,30
abuse,30
always,30
creative,30
beautiful,30
online,30
night,30
reports,29
cannot,29
choose,29
name,29
receiving,29
trademark,29
but,29
company,29
team,29
today,29
change,29
member,29
analyst,29
adding,29
valuable,29
whether,29
day,29
target,29
compare,29
plan,28
quick,28
along,28
solid,28
want,28
program,27
advice,27
priced,27
would,27
error,27
ready,27
entrepreneurial,27
ive,27
before,27
what,27
profitability,27
one,27
codes,27
available,27
contact,27
endorsement,26
legal,26
website,26
href,26
you'll,26
some,26
started,26
expires,26
seeing,26
buy,26
important,26
independent,25
research,25
quail,25
multiply,25
budget,25
historic,25
believed,25
valued,25
quality,25
having,25
trader,24
think!,24
together,24
small,24
section,24
vehicle,24
cover,24
ebooks,24
deal,24
around,24
allows,24
sharp,24
imagined,24
submitted,24
assess,24
nutshell,24
dress,23
pill,23
sign,23
active,23
unlimited,23
tough,23
campaign,23
notifying,23
believing,23
print,23
doesn't,23
backyard,23
growth,23
challenges,23
```