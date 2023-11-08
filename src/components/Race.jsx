import React from 'react'
import {
    Box,
    Card,
    CardContent,
    Typography,
    Stack,
    Chip,
    Button,
    Paper
} from '@mui/material'
import Grid from '@mui/material/Unstable_Grid2'

function CountdownTimer({ hours, minutes, seconds }) {
    return (
        <Box sx={{ flex: 1 }}>
            <Chip label={hours} />:  <Chip label={minutes} /> : <Chip label={seconds} />后结束
        </Box>
    )
}

function RaceHead({ startTime, endTime, title }) {
    return (
        <Card>
            <CardContent>
                <Typography variant="h1" component="div">
                    {title}
                </Typography>
                <Typography variant="h5" component="div">
                    {`${startTime} ~ ${endTime}`}
                </Typography>
                <Grid container spacing={2}>
                    <Grid xs={4}>
                        <Button variant="contained" sx={{ mr: 5 }}>报名</Button>
                        <Typography variant="body">
                            主办方：StarOJ
                        </Typography>
                    </Grid>
                    <Grid xs={8} sx={{ flexDirection: 'end' }}>
                        <CountdownTimer hours="17" minutes="08" seconds="34" />
                    </Grid>
                </Grid>
            </CardContent>
        </Card>
    )
}

function RaceInfo({ startTime, endTime, duration, type, problemCount }) {
    return (
        <Card>
            <CardContent>
                <Stack direction="column" spacing={3}>
                <Stack direction="row" spacing={2}>
                    <Chip label="开始时间" /> <Typography>{startTime}</Typography>
                </Stack>
                <Stack direction="row" spacing={2}>
                    <Chip label="结束时间" /> <Typography>{endTime}</Typography>
                </Stack>
                <Stack direction="row" spacing={2}>
                    <Chip label="比赛时长" /> <Typography>{duration}</Typography>
                </Stack>
                <Stack direction="row" spacing={2}>
                    <Chip label="比赛类型" /> <Typography>{type}</Typography>
                </Stack>
                <Stack direction="row" spacing={2}>
                    <Chip label="题数" /> <Typography>{problemCount}</Typography>
                </Stack>
                </Stack>
                
            </CardContent>
        </Card>
    )
}

function RaceBody() {

    return (
        <Paper elevation={3} sx={{minHeight:"70vh"}}>
            Lorem ipsum dolor, sit amet consectetur adipisicing elit. Praesentium quam iusto, quo saepe sequi nesciunt optio fugit ipsam quos nemo id. Doloribus accusantium labore reprehenderit, distinctio tempore autem voluptate quod?
        </Paper>
    )
}
function Race() {
    const title = 'StarOJ个赛'
    const startTime = '2023-09-15 15:00'
    const endTime = '2023-09-15 18:00'
    const duration = '3H'
    const type = 'IOI'
    const problemCount = 4

    return (
        <Grid container spacing={3}>
            <Grid xs={12}>
                <RaceHead startTime={startTime} endTime={endTime} title={title} />
            </Grid>
            <Grid xs={4}>
                <RaceInfo startTime={startTime} endTime={endTime} duration={duration} type={type} problemCount={problemCount} />
            </Grid>
            <Grid xs={8}>
                <RaceBody />
            </Grid>
        </Grid>
    )
}

export default Race
