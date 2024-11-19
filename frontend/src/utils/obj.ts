export function mergeObjects(obj1: any, obj2: any): any {
  for (const key in obj2) {
    if (
      obj2.hasOwnProperty(key) &&
      (obj1.hasOwnProperty(key) || !(key in obj1))
    ) {
      if (
        typeof obj2[key] === 'object' &&
        obj2[key] !== null &&
        typeof obj1[key] === 'object' &&
        obj1[key] !== null
      ) {
        mergeObjects(obj1[key], obj2[key])
      } else {
        obj1[key] = obj2[key]
      }
    }
  }
}