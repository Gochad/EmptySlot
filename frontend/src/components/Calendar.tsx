import React from "react";
import {Scheduler} from "@aldabil/react-scheduler";
import {EVENTS} from "./events";

const Calendar = () => {
    return (
        <div style={{maxWidth: '500px', maxHeight: '100px', display: 'flex', width: '300px'}}>
            <Scheduler
                events={EVENTS}
                height={10}
            />
        </div>
    );
}
export default Calendar;