export function isSameDate(date1: Date, date2: Date) {
    if (
        date1.getFullYear() === date2.getFullYear() &&
        date1.getMonth() === date2.getMonth() &&
        date1.getDate() === date2.getDate()
    ) {
        return true;
    }

    return false;
}

function zeroLeadDate(num: number) {
    return ("0" + num).slice(-2)
}

export function dateToISODateStr(year: number, month: number, day: number) {
    const zeroLeadMonth = zeroLeadDate(month + 1);
    const zeroLeadDay = zeroLeadDate(day)
    return `${year}-${zeroLeadMonth}-${zeroLeadDay}`;
}