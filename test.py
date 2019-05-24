import urllib
import urllib2

user_agent = 'Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0)'
header={'User-Agent' : user_agent}
url = "https://docs.google.com/forms/d/e/1FAIpQLScVuLBIVh-zwlbyUE9PouzxUWd_SpNT_r4Kd1TbRCvM3fSBMg/viewform"
# values from your form. You will need to include any hidden variables if you want to..
values= {
'entry.asdfsdfsdasd': 'asdfasdfsd',
'draftResponse':'[,,&quot;-asdfasdasdf&quot;]',
'pageHistory':'0',
'fbzx':'-asdfasdfsd'
}
data = urllib.urlencode(values)
urllib2.Request(url, data, header)
