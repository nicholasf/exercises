export interface CacheValue {
    value: number,
    when: number
}

export class LRUCache {
    members: Map<number, CacheValue>
    capacity: number

    constructor(capacity: number) {
        this.members = new Map<number, CacheValue>
        this.capacity = capacity
    }

    get(key: number): number {
        let val = this.members.get(key);

        if (val == null) {
            return -1
        }        

        val.when = performance.now();

        return val.value;
    }

    put(key: number, value: number): void {
        this.members.set(key,{ value: value, when: performance.now() });

        if (this.members.size > this.capacity) {
            this.evict()
        } 
    }

    inspect(): void {
        // console.log('cache:');
        this.members.forEach((value, key): void => {
            // console.log(`  ${key}=${value.value} (${value.when})`)
        })
    }   

    evict(): void {
        let candidateValue: CacheValue;
        let candidateKey: number;

        this.members.forEach((value, key): void => {
            if (candidateValue == null) {
                candidateValue = value;
                candidateKey = key; 
            }

            // console.log(`candidate: ${candidateKey}=${candidateValue.value} (${candidateValue.when}), current: ${key}=${value.value} (${value.when})`)

            if (candidateValue.when > value.when) {
                candidateKey = key;
                candidateValue = value;
                // console.log(`new candidate: ${candidateKey}=${candidateValue.value} (${candidateValue.when})`);
            }
        })

        // console.log(`evicting ${candidateKey!}=${candidateValue!.value}`);
        this.members.delete(candidateKey!);
    }
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * var obj = new LRUCache(capacity)
 * var param_1 = obj.get(key)
 * obj.put(key,value)
 */