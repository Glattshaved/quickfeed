export class MapHelper {
    public static mapTo<T, T2>(map: IMap<T>, callback: (ele: T, index: number, map: IMap<T>) => T2): T2[] {
        const returnArray: T2[] = [];
        const keys = Object.keys(map);
        for (const a of keys) {
            const index = parseInt(a, 10);
            const temp = map[index];
            if (temp) {
                returnArray.push(callback(temp, index, map));
            }
        }
        return returnArray;
    }

    public static forEach<T>(map: IMap<T>, callback: (ele: T, index: number, map: IMap<T>) => void): void {
        const keys = Object.keys(map);
        for (const a of keys) {
            const index = parseInt(a, 10);
            const temp = map[index];
            if (temp) {
                callback(temp, index, map);
            }
        }
    }

    public static find<T>(map: IMap<T>, callback: (ele: T, index: number, map: IMap<T>) => boolean): T | null {
        const keys = Object.keys(map);
        for (const a of keys) {
            const index = parseInt(a, 10);
            const temp = map[index];
            if (temp && callback(temp, index, map)) {
                return temp;
            }
        }
        return null;
    }

    public static toArray<T>(map: IMap<T>): T[] {
        const returnArray: T[] = [];
        const keys = Object.keys(map);
        for (const a of keys) {
            const index = parseInt(a, 10);
            const temp = map[index];
            if (temp) {
                returnArray.push(temp);
            }
        }
        return returnArray;
    }
}

export interface IMap<T> {
    [index: number]: T | undefined;
}

export function mapify<T>(obj: T[], callback: (ele: T, index: number, array: T[]) => number) {
    const newObj: IMap<T> = {};
    obj.forEach((ele, index, array) => {
        newObj[callback(ele, index, obj)] = ele;
    });
    return newObj;
}
