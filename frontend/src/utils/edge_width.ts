export function calcEdgeWidthFromPathBetweenness(
  betweenness: number,
  size: number = 1,
) {
  return Math.max(betweenness * size * 2, 0.5)
}
