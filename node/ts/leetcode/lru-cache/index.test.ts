import { describe, expect, test } from '@jest/globals';
import { CacheValue, LRUCache } from './index.js';

describe('LRUCache', () => {
  test('stores and retrieves values', () => {
        let cache = new LRUCache(2);
        cache.put(1, 1);
        cache.put(2, 2);

        expect(cache.get(1)).toBe(1);
        expect(cache.get(2)).toBe(2);
    })

  test('returns -1 for missing keys', () => {
        let cache = new LRUCache(2);
        expect(cache.get(1)).toBe(-1);
    })

  test('get updates the when of the cache entry', () => {
        let cache = new LRUCache(2);
        cache.put(1, 1);
        cache.put(2, 2);

        cache.get(2); // this updates the when of 2, making it more recent than 1

        let val1 = cache.members.get(1);
        let val2 = cache.members.get(2);

        expect(val2!.when).not.toBe(val1!.when);
        expect(val2!.when).toBeGreaterThan(val1!.when);
      });

    test('evicts least recently used keys', () => {
        let cache = new LRUCache(2);
        cache.put(1, 1);
        cache.put(2, 2);
        cache.put(3, 3);

        expect(cache.get(1)).toBe(-1);
    })

    test('leetcode example;', () => {
      let lRUCache: LRUCache = new LRUCache(2);
      lRUCache.put(1, 1); // cache is {1=1}
      lRUCache.put(2, 2); // cache is {1=1, 2=2}
      expect(lRUCache.get(1)).toBe(1);    // return 1
      lRUCache.inspect();
      lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
      expect(lRUCache.get(2)).toBe(-1);    // returns -1 (not found)
      lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
      expect(lRUCache.get(1)).toBe(-1);    // return -1 (not found)
      expect(lRUCache.get(3)).toBe(3);    // return 3
      expect(lRUCache.get(4)).toBe(4);    // return 4
  });
});

