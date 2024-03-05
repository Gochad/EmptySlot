import React, { useState } from "react";
import { Calendar, momentLocalizer } from "react-big-calendar";
import moment from "moment";
import "react-big-calendar/lib/css/react-big-calendar.css";

moment.locale("en-GB");
const localizer = momentLocalizer(moment);

export default function FullCalendar() {
    const [eventsData, setEventsData] = useState([]);

    const handleSelect = ({ start, end }) => {
        console.log(start);
        console.log(end);
        const title = window.prompt("New Event name");
        if (title)
            setEventsData([
                ...eventsData,
                {
                    start,
                    end,
                    title
                }
            ]);
    };

    return (
        <div className="App">
            <Calendar
                localizer={localizer}
                events={eventsData}
                startAccessor="start"
                endAccessor="end"
                style={{ height: 500 }}
                selectable
                onSelectSlot={handleSelect}
            />
        </div>
    );
}