import React from 'react';
import { Scheduler } from "@aldabil/react-scheduler";
import { EVENTS } from "./events";



function App() {
    return (
        <div style={{ maxWidth: '300x', maxHeight: '10px', display: 'flex' }}>
            <Scheduler
                events={EVENTS}
                height={10}
                // week={{
                //   weekDays: [0, 1, 2, 3, 4, 5, 6],
                //   weekStartOn: 6,
                //   startHour: 0,
                //   endHour: 24,
                //   step: 30
                // }}
            />
        </div>

    );
}

export default App;
