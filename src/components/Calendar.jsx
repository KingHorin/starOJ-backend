import React from 'react';
import Paper from '@mui/material/Paper';
import Grid from '@mui/material/Unstable_Grid2';
import Tooltip from '@mui/material/Tooltip';

const generateRandomData = () => {
    const data = {};
    const startDate = new Date();
    startDate.setMonth(startDate.getMonth() - 3);

    for (let i = 0; i < 90; i++) {
        const currentDate = new Date(startDate);
        currentDate.setDate(currentDate.getDate() + i);
        const dateStr = currentDate.toISOString().split('T')[0];
        data[dateStr] = Math.floor(Math.random() * 10); // 随机生成 0 到 9 的完成数量
    }

    return data;
};

const Calendar = () => {
    const data = generateRandomData();

    const calendarData = [];

    let currentDate = new Date();
    currentDate.setMonth(currentDate.getMonth() - 3);
    const endDate = new Date();

    while (currentDate <= endDate) {
        const dateStr = currentDate.toISOString().split('T')[0];
        const completed = data[dateStr] || 0;
        calendarData.push({ date: dateStr, completed });
        currentDate.setDate(currentDate.getDate() + 1);
    }

    const maxCompleted = Math.max(...calendarData.map((item) => item.completed));

    // 创建一个二维数组来表示格子排列
    const columns = [];
    let currentColumn = [];
    calendarData.forEach((grid, index) => {
        currentColumn.push(
            <Grid key={grid.date}>
                <Tooltip title={`${grid.date} 完成题目数量: ${grid.completed}`} sx={{padding:0, margin:0}} arrow>
                    <Paper
                        square
                        sx={{
                            height: '3px',
                            width: '3px',
                            backgroundColor: `rgba(0, 128, 0, ${grid.completed / maxCompleted})`,
                        }}
                    >
                    </Paper>
                </Tooltip>
            </Grid>
        );

        if (currentColumn.length === 7 || index === calendarData.length - 1) {
            columns.push(
                <Grid key={index}>
                    <Grid container sx={{
                        flexDirection: 'column', minWidth: '3px', minheight: '40px', justifyContent: 'space-between'
                    }}>
                        {currentColumn}
                    </Grid>
                </Grid>
            );
            currentColumn = [];
        }
    });

    return (
        <Grid container sx={{
            flexDirection: 'row',
            justifyContent: 'space-between',
            width: '50px',
            height: '40px'
        }}>
            {columns}
        </Grid>
    );
};

export default Calendar;
