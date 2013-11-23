# -*- coding: utf-8 -*-
from __future__ import print_function, division, absolute_import, unicode_literals

import sys
import json
from cStringIO import StringIO

from rivescript import RiveScript


def parse_input():
    """Read JSON from input."""
    request = StringIO()
    while True:
        line = raw_input()
        if line == '__END__':
            break
        request.write(line)
    raw_json = request.getvalue()
    return json.loads(raw_json)


def process_reply(reply):
    if reply.startswith('ERR: '):
        return {'status': 'error', 'reply': reply[5:]}
    return {'status': 'ok', 'reply': reply}


if __name__ == '__main__':
    rs = RiveScript()
    rs.load_directory('./brain')
    rs.sort_replies()

    while True:
        response = {}
        try:
            request = parse_input()
            message = request['message']
        except (ValueError, TypeError):
            response['status'] = 'error'
            response['reply'] = 'Failed to decode incoming JSON data.'
        except KeyError:
            response['status'] = 'error'
            response['reply'] = 'Request must contain key "message".'
        except (KeyboardInterrupt, EOFError):
            sys.exit(0)
        else:
            reply = rs.reply('localuser', message)
            response = process_reply(reply)
        print(json.dumps(response))
