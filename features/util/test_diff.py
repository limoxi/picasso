from features.util.bdd_util import diff, assert_json

if __name__ == "__main__":
    e = ['a', 'c', 1]
    a = ['b', 'c']

    assert_json(e, a)