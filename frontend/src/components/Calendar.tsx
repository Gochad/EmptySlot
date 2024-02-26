import React, {useState} from "react";
import {Scheduler} from "@aldabil/react-scheduler";

const Calendar = () => {
    const [events, setEvents] = useState([]);

    const translations = {
        navigation: {
            month: "Month",
                week: "Week",
                day: "Day",
                today: "Today",
            agenda: "Agenda"
        },
        form: {
            addTitle: "Add merchandise",
                editTitle: "Edit merchandise",
                confirm: "Confirm",
                delete: "Delete",
                cancel: "Cancel"
        },
        event: {
            title: "Title",
                start: "Start",
                end: "End",
                allDay: "All Day"
        },
        moreEvents: "More...",
        loading: "Loading..."
    }



    return (
        <div style={{maxWidth: '500px', maxHeight: '100px', display: 'flex', width: '350px'}}>
            <Scheduler
                events={events}
                height={10}
                translations={translations}
            />
        </div>
    );
}
export default Calendar;