import React from 'react';
import { Scheduler } from "@aldabil/react-scheduler";
import { EVENTS } from "./events";



function App() {
    return (
        <div style={{ maxWidth: '300x', maxHeight: '10px', display: 'flex' }}>
            <Scheduler
                events={EVENTS}
                height={10}
            />
        </div>

    );
}

export default App;
