export function calcEdgeWidthFromPathBetweenness(
  betweenness: number,
  size: number = 1,
) {
  return betweenness * size * 2 + 1
}
