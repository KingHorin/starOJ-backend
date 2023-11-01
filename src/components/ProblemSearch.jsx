import React, { useState, useEffect } from 'react'
import {
  FormLabel,
  TextField,
  RadioGroup,
  Radio,
  FormControlLabel,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  Table,
  TableHead,
  TableRow,
  TableCell,
  TableBody,
  Chip,
  LinearProgress
} from '@mui/material'
import Grid from '@mui/material/Unstable_Grid2'

const SearchBox = ({ filters, setFilters, problems, setProblems }) => {
  const handleInputChange = e => {
    setFilters(prev => ({ ...prev, name: e.target.value }))
  }

  const handleProblemSetChange = e => {
    setFilters(prev => ({ ...prev, problemSet: e.target.value }))
  }
  const handleDifficultyChange = e => {
    setFilters(prev => ({ ...prev, difficulty: e.target.value }))
  }

  const handleTagsChange = e => {
    setFilters(prev => ({ ...prev, tags: e.target.value }))
  }

  const handleYearChange = e => {
    setFilters(prev => ({ ...prev, authoringYear: e.target.value }))
  }

  return (
    <Grid container spacing={3} justifyContent='center'>
      <Grid xs={8}>
        <TextField
          fullWidth
          label='输入题目名称查询'
          onChange={handleInputChange}
        />
      </Grid>
      <Grid xs={4}>
        <TextField
          fullWidth
          type='number'
          label='时间'
          onChange={handleYearChange}
        />
      </Grid>
      <Grid xs={8}>
        <FormControl fullWidth>
          <FormLabel>题库</FormLabel>
          <RadioGroup row onChange={handleProblemSetChange}>
            <FormControlLabel
              value='StarOJ'
              control={<Radio />}
              label='StarOJ'
            />
            <FormControlLabel value='其它' control={<Radio />} label='其它' />
          </RadioGroup>
        </FormControl>
      </Grid>
      <Grid xs={8}>
        <FormControl fullWidth>
          <FormLabel>难度</FormLabel>
          <RadioGroup row onChange={handleDifficultyChange}>
            <FormControlLabel value='全部' control={<Radio />} label='全部' />
            <FormControlLabel value='简单' control={<Radio />} label='简单' />
            <FormControlLabel value='中等' control={<Radio />} label='中等' />
            <FormControlLabel value='困难' control={<Radio />} label='困难' />
          </RadioGroup>
        </FormControl>
      </Grid>
      <Grid xs={8}>
        <FormControl fullWidth>
          <InputLabel>标签</InputLabel>
          <Select
            multiple
            value={filters.tags}
            onChange={handleTagsChange}
            renderValue={selected => (
              <div>
                {selected.map(value => (
                  <Chip key={value} label={value} style={{ margin: '2px' }} />
                ))}
              </div>
            )}
          >
            <MenuItem value='算法'>算法</MenuItem>
            <MenuItem value='数学'>数学</MenuItem>
            <MenuItem value='动态规划'>动态规划</MenuItem>
            <MenuItem value='树'>树</MenuItem>
            <MenuItem value='字符串'>字符串</MenuItem>
          </Select>
        </FormControl>
      </Grid>
    </Grid>
  )
}

const SearchResults = ({ problems, filters }) => {
  const filteredProblems = problems.filter(problem => {
    let matches = true

    if (filters.name && !problem.name.includes(filters.name)) matches = false
    if (
      filters.difficulty !== '全部' &&
      problem.difficulty !== filters.difficulty
    )
      matches = false
    if (filters.tags && !filters.tags.every(tag => problem.tags.includes(tag)))
      matches = false
    if (
      filters.authoringYear &&
      problem.authoringYear !== filters.authoringYear
    )
      matches = false
    if (filters.problemSet && filters.problemSet !== problem.problemSet)
      matches = false
    return matches
  })

  return (
    <Table>
      <TableHead>
        <TableRow>
          <TableCell>状态</TableCell>
          <TableCell style={{ width: '50%' }}>题目</TableCell>
          <TableCell>标签</TableCell>
          <TableCell>难度</TableCell>
          <TableCell>通过率</TableCell>
        </TableRow>
      </TableHead>
      <TableBody>
        {filteredProblems.map(problem => (
          <TableRow key={problem.name}>
            <TableCell>
              <Chip label={problem.state} style={{ marginRight: '5px' }} />
            </TableCell>
            <TableCell>{problem.name}</TableCell>
            <TableCell>
              {problem.tags.map(tag => (
                <Chip key={tag} label={tag} style={{ marginRight: '5px' }} />
              ))}
            </TableCell>
            <TableCell>
              <Chip label={problem.difficulty} style={{ marginRight: '5px' }} />
            </TableCell>
            <TableCell>
              <LinearProgress
                variant='determinate'
                value={parseInt(problem.passRate, 10)}
              />
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  )
}

const ProblemSearch = () => {
  const [problems, setProblems] = useState([])
  const [filters, setFilters] = useState({
    authoringYear: undefined,
    name: '',
    difficulty: '全部',
    tags: [],
    problemSet: 'StarOJ'
  })

  useEffect(() => {
    const randomProblems = Array.from({ length: 50 }).map((_, index) => {
      const randomDifficulty = ['简单', '中等', '困难'][
        Math.floor(Math.random() * 3)
      ]
      const randomTags = ['算法', '数学', '动态规划', '树', '字符串'].filter(
        () => Math.random() > 0.5
      )
      const randomYear = Math.floor(Math.random() * (2023 - 2015 + 1)) + 2015 // Random year between 2015 and 2023
      const randomPassRate = Math.floor(Math.random() * 101) // 0 to 100
      const problemSet = Math.random() > 0.5 ? 'StarOJ' : '其它'

      return {
        state: Math.random() > 0.5 ? '已通过' : '未通过',
        authoringYear: randomYear,
        name: `Problem ${index + 1}`,
        difficulty: randomDifficulty,
        tags: randomTags,
        passRate: `${randomPassRate}%`,
        problemSet
      }
    })

    setProblems(randomProblems)
  }, [])

  return (
    <div>
      <SearchBox
        filters={filters}
        setFilters={setFilters}
        problems={problems}
        setProblems={setProblems}
      />
      <SearchResults problems={problems} filters={filters} />
    </div>
  )
}

export default ProblemSearch
