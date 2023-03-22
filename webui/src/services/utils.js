var utils = {}

/*
    *  Convert a timestamp to a string representing the time elapsed since the timestamp.
    *  @param {Number} timestamp - The timestamp to convert.
    * @returns {String} The time elapsed since the timestamp.
*/
utils.since = (timestamp) => {
    const ONE_MINUTE = 60 * 1000;
    const ONE_HOUR = 60 * ONE_MINUTE;
    const ONE_DAY = 24 * ONE_HOUR;
    const ONE_WEEK = 7 * ONE_DAY;

    timestamp = new Date(timestamp);

    const diff = Math.abs(Date.now() - timestamp);
    const elapsedMinutes = Math.floor(diff / ONE_MINUTE);
    const elapsedHours = Math.floor(diff / ONE_HOUR);
    const elapsedDays = Math.floor(diff / ONE_DAY);
    const elapsedWeeks = Math.floor(diff / ONE_WEEK);

    if (elapsedWeeks > 0) {
        return `${elapsedWeeks}w`;
    } else if (elapsedDays > 0) {
        return `${elapsedDays}d`;
    } else if (elapsedHours > 0) {
        return `${elapsedHours}h`;
    } else {
        return `${elapsedMinutes}m`;
    }
}

/*
    *  Convert an error to a string. If the error has a response, the response data is returned.
    *  Otherwise, the error is converted to a string.
    *  @param {Error} error - The error to convert.
    *  @returns {String} The string representation of the error.
*/
utils.errorToString = (error) => {
    console.log(error);
    if (error.hasOwnProperty('response')) {
        return error.response.data;
    } else {
        return error.toString();
    }
}

export default utils;