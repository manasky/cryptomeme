const epochsShort = [
  ['y', 31536000],
  ['mo', 2592000],
  ['d', 86400],
  ['h', 3600],
  ['m', 60],
  ['s', 1]
];
const epochsLong = [
  ['year', 31536000],
  ['month', 2592000],
  ['day', 86400],
  ['hour', 3600],
  ['minute', 60],
  ['second', 1]
];

const getDuration = (timeAgoInSeconds, epochs) => {
  for (let [name, seconds] of epochs) {
      const interval = Math.floor(timeAgoInSeconds / seconds);
      if (interval >= 1) {
          return {
              interval: interval,
              epoch: name
          };
      }
  }
};

export const timeAgo = (date, fullNames=false) => {
  let timeAgoInSeconds = Math.floor((new Date() - new Date(date)) / 1000);
  let {interval, epoch} = getDuration(timeAgoInSeconds, epochsShort);
  if (fullNames) {
    let {interval, epoch} = getDuration(timeAgoInSeconds, epochsLong);
    let suffix = interval === 1 ? '' : 's';
    return `${interval} ${epoch}${suffix}`;
  }
  return `${interval}${epoch}`;
};