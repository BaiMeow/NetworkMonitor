export function parseDuration(duration: string) {
  let seconds = 0
  const hours = duration.match(/(\d+)h/)
  const minutes = duration.match(/(\d+)m/)
  const secs = duration.match(/(\d+)s/)
  if (hours) seconds += parseInt(hours[1]) * 3600
  if (minutes) seconds += parseInt(minutes[1]) * 60
  if (secs) seconds += parseInt(secs[1])
  return seconds
}
