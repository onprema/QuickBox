#!/usr/bin/python3
'''
This module holds a dictionary that acts as the database for container images.
'''
import json
import uuid

class Container(object):
    '''Defines the container.'''

    def __init__(self, base):
        '''Initializes container with id and base.'''
        self.id = str(uuid.uuid4())
        self.base = base

    def to_json(self):
        '''Returns a jsonified version of the container.'''
        filename = '{}.json'.format(self.id)
        with open(filename, 'w') as f:
            json.dump(self.__dict__, f)
