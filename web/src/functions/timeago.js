const epochs = [
  ['y', 31536000],
  ['mo', 2592000],
  ['d', 86400],
  ['h', 3600],
  ['m', 60],
  ['s', 1]
];

const getDuration = (timeAgoInSeconds) => {
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

export const timeAgo = (date) => {
  const timeAgoInSeconds = Math.floor((new Date() - new Date(date)) / 1000);
  const {interval, epoch} = getDuration(timeAgoInSeconds);
  const suffix = interval === 1 ? '' : 's';
  return `${interval}${epoch}`;
};