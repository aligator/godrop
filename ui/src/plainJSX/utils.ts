let lastRefId = 0

type Ref<T> = {
    id: string
    update: (cb: (ref: T) => void) => void
}

/**
 * Allows you to bind a HTMLElement to a variable.
 * Just assign the id to the element you want to bind.
 * 
 * To use it, just call update(cb) on the reference and
 * you can use it inside the callback.
 * The callback will only be called when the id could be found.
 * 
 * @returns the reference object.
 */
export const useRef = <T extends HTMLElement>(): Ref<T> => {
    let ref: T | undefined = undefined
    
    const res: Ref<T> = {
        id: `ref-${lastRefId++}`,
        update: (cb) => {
            if (ref === undefined) {
                const element = document.getElementById(res.id)
                if (element) {
                    ref = element as T
                }
            }

            if (ref === undefined) {
                return
            }

            cb(ref)
        }
    }

    return res
}

export interface Listener<T> {
    (currentValue: T): void;
    readonly id: number;
    unsubscribe(): void
}

interface IdWriteableListener<T> {
    (currentValue: T): void;
    id: number;
    unsubscribe(): void
}

export class Observable<T> {
    private lastListenerId = 0
    private readonly listeners: Map<number, Listener<T>> = new Map()

    private backingValue: T
    set value(newValue: T) {
        this.backingValue = newValue
        this.update(this.backingValue)
    }
    get value(): T {
        return this.backingValue
    }

    constructor(defaultValue: T) {
        this.backingValue = defaultValue
        this.update(this.backingValue)
    }

    unsubscribe(id: number): void {
        this.listeners.delete(id)
    }

    listen(listener: (currentValue: T) => void): Listener<T> {
        this.lastListenerId++
        const newId = this.lastListenerId

        const newListener: Listener<T> = (() => {
            const _l = listener as Listener<T>
            // Force here the setting of the id value to be able
            // to set it one time. After that the value should be fixed.
            (_l as IdWriteableListener<T>).id = newId;
            _l.unsubscribe = () => this.unsubscribe(_l.id)
            return _l;
        })();

        this.listeners.set(newId, newListener)

        return newListener
    }

    protected update(newValue: T): void {
        this.listeners.forEach((l) => l(newValue))
    }

    destroy(): void {
        this.listeners.forEach((l, k) => {
            this.listeners.delete(k)
        })
    }
}
