import plotly.graph_objects as go

logPart = """
3	loading	0.184901	0.209007
2	storage	0.205416	0.227645
4	loading	0.209011	0.246545
3	parsing	0.209074	0.228591
3	storage	0.228596	0.237616
5	loading	0.246548	0.281826
4	parsing	0.246586	0.263061
4	storage	0.263066	0.282869
"""

blockData = []
for line in logPart.strip().split('\n'):
        id, actType, start, end = line.split('\t')
        if actType not in ['loading', 'parsing', 'storage']:
            continue
        blockData.append((int(id), actType, float(start), float(end)))

actYRange = {
    'loading': (2, 3),
    'parsing': (1, 2),
    'storage': (0, 1),
}

actColor = {
    'loading': 'blue',
    'parsing':'red',
    'storage': 'green',
}


# go.Scatter(x=[0,1,2,0,None,3,3,5,5,3], y=[0,2,0,0,None,0.5,1.5,1.5,0.5,0.5], fill="toself")

fig = go.Figure()
for id, actType, start, end in blockData:
    yLow = actYRange[actType][0]
    yHigh = actYRange[actType][1]
    fig.add_trace(go.Scatter(x=[start, start, end, end, start], y=[yLow, yHigh, yHigh, yLow, yLow], line=go.Line({"color":'black'}), fill='toself', fillcolor=actColor[actType], showlegend=False))
fig.update_yaxes(tickvals=[0.5, 1.5, 2.5], ticktext=["storage", "parsing", "loading"], tickfont=dict(size=20))
fig.update_yaxes(title="Вид обработчика", titlefont=dict(size=20))
fig.update_xaxes(title="Время, секунды", titlefont=dict(size=20), tickfont=dict(size=20))
fig.show()
